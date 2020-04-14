## 日志组件
对 go.uber.org/zap 的封装，按照时间和日志类型分割

## usage
```go
package main

import "github.com/strconv/valyria/log"

func init() {
    // info、debug、error
	log.InitLog("info")
}

func main() {
	log.Errorf("an error log |err:%s|", "have fun! ")
}

```

## todo
- debug 单独输出
- 打印频率 day、hour
- 添加 trace
- ...

## 参考
https://www.liwenzhou.com/posts/Go/zap/
https://www.cnblogs.com/Me1onRind/p/10918863.html
https://studygolang.com/articles/25044?fr=sidebar