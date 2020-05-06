## 配置文件加载
特性：加载配置文件、自定义配置字段

## usage
```go
package main

import (
	"fmt"

	"github.com/strconv/valyria/config"
)

var Extra struct {
    Token  string `yaml:"token"`
    AppKey string `yaml:"appKey"`
}

func init() {
    config.Init("./conf.yaml", &Extra)
}

func main() { 
    // 公共配置信息 
    fmt.Printf("service_name: %s", config.Conf.Service.Name)
    // 自定义配置信息
    fmt.Println("token: ", Extra.Token, "app_key: ", Extra.AppKey)
}

```

```yaml
# 服务信息
service:
  name: user.account
  addr: :10086
  logLevel: info # 日志等级

# consul 配置
consul: 127.0.0.1:8500

# jaeger 配置
jaeger:  127.0.0.1:6831

# mysql 
# 非必填（若填入，需要有mysql环境支持）
mysql:
  addr: 127.0.0.1
  port : 3306
  dbName : rbac
  username : root
  password : gogocuri

# redis 配置
redis:
  host: 127.0.0.1:6890
  auth: abc123
  db:  0
  maxIdle: 200 # 最大空闲数
  idleTimeout: 2000 # 超时时间
  maxActive: 400 # 连接池最大连接数

# JWT
jwt:
  secret: "xxxxxxx"
  timeout: 120 # 分钟

# 自定义
extra:
  appKey: sdasdsa
  token: token_xxzcdasxx
```

## todo
暂时支持用 yaml 吧