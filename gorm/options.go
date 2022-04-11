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

package gorm

type DBType string

const (
	// UNKNOWN unknown database
	UNKNOWN DBType = "unknown"
	// MYSQL mysql
	MYSQL DBType = "sql"
	// IPV4 others database type
	IPV4 DBType = "others"
)

const (
	componentIDUnknown  = 0
	componentIDMysql    = 5012
	componentIDOtherSql = 5999
)

type Option func(*options)

type options struct {
	dbType      DBType
	peer        string
	componentID int32
	dbName      string
	reportQuery bool
	reportParam bool
}

// WithSqlDBType set dbType option,
// dbType is used for parsing dsn string to peer address
// and setting componentID, if DB type is not support in DBType
// list, please use WithPeerAddr to set peer address manually
func WithSqlDBType(t DBType) Option {
	return func(o *options) {
		//2022-02-26 huangyao 由于oap端进行sql统计和明细处理（类：MultiScopesSpanListener）中判断为"sql"所以此处定义为sql
		// o.dbType = t
		o.setDbType(t)
		o.setComponentID()
	}
}

// WithPeerAddr set the peer address to report
func WithPeerAddr(addr string) Option {
	return func(o *options) {
		o.peer = addr
	}
}

func WithQueryReport() Option {
	return func(o *options) {
		o.reportQuery = true
	}
}

func WithParamReport() Option {
	return func(o *options) {
		o.reportParam = true
	}
}

func (o *options) setComponentID() {
	switch o.dbType {
	case MYSQL:
		o.componentID = componentIDMysql
	//2022-02-26 黄尧 非Mysql的传递componentID为5999，在oap的component-libraries.yml中定义
	default:
		o.componentID = componentIDOtherSql
	}
}

func (o *options) setDbType(dbType DBType) {
	switch dbType {
	case MYSQL:
		o.dbType = MYSQL
	default:
		o.dbType = "sql"
	}
}
