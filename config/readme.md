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

# mysql
mysql:
  addr: 127.0.0.1
  port : 3306
  db_name : rbac
  username : root
  password : gogocuri

# redis 配置
redis:
  host: 127.0.0.1
  port: 6379
  auth: abc123
  db:  0

# consul 配置
consul: 127.0.0.1:8500

# jaeger 配置
jaeger:  127.0.0.1:6831

# 自定义
extra:
  appKey: sdasdsa
  token: token_xxzcdasxx
```

## todo
暂时支持用 yaml 吧