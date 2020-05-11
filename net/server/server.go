package server

import (
	"time"

	"github.com/strconv/valyria/log"

	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/strconv/valyria/config"
	"github.com/strconv/valyria/middleware"
	"github.com/strconv/valyria/trace"
)

const (
	REGISTER_TTL      = 15 //重新注册时间
	REGISTER_INTERVAL = 10 //注册过期时间
)

func NewHTTP() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// 公共中间件
	// trace
	router.Use(middleware.TracerWrapper)
	router.Use(log.GinLogger())
	return router
}

func InitHTTP(conf *config.Conf, handler *gin.Engine) {
	// consul
	reg := consul.NewRegistry(
		registry.Addrs(conf.Consul),
	)

	service := web.NewService(
		web.Name(conf.Service.Name),
		web.Address(conf.Service.Addr),
		web.RegisterTTL(time.Second*REGISTER_TTL),           // 设置注册服务的过期时间
		web.RegisterInterval(time.Second*REGISTER_INTERVAL), // 设置间隔多久再次注册服务
		web.Handler(handler),                                // use gin's handler
		web.Registry(reg),
	)

	err := service.Init(
		web.Action(func(context *cli.Context) {
			// trace
			trace.Init(50, config.C.Service.Name, config.C.Jaeger)
			// ...
		}),
	)

	if err != nil {
		panic("init micro service fail: " + err.Error())
	}

	if err := service.Run(); err != nil {
		panic("micro run fail: " + err.Error())
	}
}
