//
// Copyright 2022 SkyAPM org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gear

import (
	"fmt"
	"strconv"
	"time"

	"github.com/powerapm/go2sky"
	"github.com/powerapm/go2sky/propagation"
	commonv2 "github.com/powerapm/go2sky/reporter/grpc/common"
	"github.com/teambition/gear"
)

const componentIDGearServer = 5007

//Middleware gear middleware return HandlerFunc  with tracing.
func Middleware(tracer *go2sky.Tracer) gear.Middleware {
	return func(ctx *gear.Context) error {
		if tracer == nil {
			return nil
		}
		//2022-04-06 黄尧 创建entrySpan时传递ctx.Context()作为上下文信息传递，
		//使得传入的类型统一为context.Context,方便结尾出ctx.WithContext(nCtx)的赋值（否则会报错）
		// span, nCtx, err := tracer.CreateEntrySpan(ctx, operationName(ctx), func() (string, error) {
		span, nCtx, err := tracer.CreateEntrySpan(ctx.Context(), operationName(ctx), func() (string, error) {
			return ctx.GetHeader(propagation.Header), nil
		})
		if err != nil {
			return nil
		}

		span.SetComponent(componentIDGearServer)
		span.Tag(go2sky.TagHTTPMethod, ctx.Method)
		span.Tag(go2sky.TagURL, ctx.Host+ctx.Path)
		span.SetSpanLayer(commonv2.SpanLayer_Http)

		ctx.OnEnd(func() {
			code := ctx.Res.Status()
			span.Tag(go2sky.TagStatusCode, strconv.Itoa(code))
			if code >= 400 {
				span.Error(time.Now(), string(ctx.Res.Body()))
			}
			span.End()
		})
		//2022-04-06 黄尧 将上下文信息存放到request和gearContext的ctx信息中，方便后续调用流程获取进行传递(将链路串起来)
		ctx.WithContext(nCtx)
		return nil
	}
}

func operationName(ctx *gear.Context) string {
	return fmt.Sprintf("/%s%s", ctx.Method, ctx.Path)
}
