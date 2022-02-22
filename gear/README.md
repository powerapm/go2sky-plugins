# Go2sky with gear (v1.21.2)

## Installation

```bash
go get -u github.com/powerapm/go2sky-plugins/gear
```

## Usage
```go
package main

import (
	"log"

	"github.com/powerapm/go2sky"
	gearplugin "github.com/powerapm/go2sky-plugins/gear"
	"github.com/powerapm/go2sky/reporter"
	"github.com/teambition/gear"
)

func main() {
    // Use gRPC reporter for production
	re, err := reporter.NewLogReporter()
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}

	defer re.Close()

	tracer, err := go2sky.NewTracer("gear", go2sky.WithReporter(re))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}

	app := gear.New()
    
	//Use go2sky middleware with tracing
	app.Use(gearplugin.Middleware(tracer))

	// do something
}
```

[See more](example_gear_test.go).