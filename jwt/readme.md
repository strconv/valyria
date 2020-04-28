## JWT封装
参考：
- [在gin框架中使用JWT](https://www.liwenzhou.com/posts/Go/jwt_in_gin/)
- 
## feature
- JWT-Token的生成和解析
- 解析错误类型分类
- 适用于 `gin` 的中间件封装（见`valyria/middleware/jwt.go`）
- 续签

## usage
### 基本使用

```go
package main

import (
	"fmt"
	"github.com/strconv/valyria/jwt"
)

func main() {
	// 生成
	token, _ := jwt.GenToken(1000005)
	
	// 解析
	c, err := jwt.ParseToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.UID)
}

```

### 中间件
对 gin 中每个需要JWT认证的接口做鉴权
```go
account := s.Group("/user/info")
// 在这里添加中间件
account.Use(middleware.JWTAuth()) // JWT
{
    account.GET("/get", userInfo.Get)
    account.POST("/update", userInfo.Update)
}
```
需要注意的是
- 如果用 `valyria` ，需要在配置文件内加上JWT相关配置
    ```yaml
      jwt:
        secret: "welcome to ariser.cn"
        timeOut: 120 # 分钟
    ```

## TODO
- JWT自定义参数
- `timeout` 和 `secret` 作为参数传入
