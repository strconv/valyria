## 日志组件
对 go.uber.org/zap 的封装，按照时间和日志类型分割

## feature
- 设置日志级别
- 集成 trace ID

## usage
```go
package main

import (
	"context"

	"github.com/strconv/valyria/log"
)

func init() {
	// info、debug、error
	log.Init("debug")
}

func main() {
	log.Debug("an debug log")
	log.For(context.Background(), "func", "log_test").Infof("an info log |msg:%s|", "have fun! ")
	log.Error("an error log")
}

```

## todo
- 打印频率 day、hour
- 添加`client-log`
- 添加`gin-log`

## 参考
https://www.liwenzhou.com/posts/Go/zap/
https://www.cnblogs.com/Me1onRind/p/10918863.html
https://studygolang.com/articles/25044?fr=sidebar