# Go2Sky with Gorm

## Installation

```bash
go get -u github.com/powerapm/go2sky-plugins/gorm
```

## Usage

```go
import (
	gormPlugin "github.com/powerapm/go2sky-plugins/gorm"

	"github.com/powerapm/go2sky"
	"github.com/powerapm/go2sky/reporter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// init reporter
re, err := reporter.NewLogReporter()
defer re.Close()

// init tracer
tracer, err := go2sky.NewTracer("service-name", go2sky.WithReporter(re))
if err != nil {
    log.Fatalf("init tracer error: %v", err)
}

db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

if err != nil {
  log.Fatalf("open db error: %v \n", err)
}
db.Use(gormPlugin.New(tracer, "127.0.0.1:3306", gormPlugin.MYSQL))

// use with context
dbWithCtx := db.WithContext(ctx)
```