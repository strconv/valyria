package server

import (
	"time"

	"github.com/strconv/valyria/middleware"

	"github.com/strconv/valyria/trace"

	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/strconv/valyria/config"
)

const (
	REGISTER_TTL      = 15 //重新注册时间
	REGISTER_INTERVAL = 10 //注册过期时间
)

func NewHTTP() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router.Use() //middleware trace、swagger、
	// 公共中间件
	router.Use(middleware.TracerWrapper) // trace 包装
	// todo:: swagger、ginLog
	return router
}

func InitHTTP(conf *config.Conf, handler *gin.Engine) {
	// consul
	reg := consul.NewRegistry(
		registry.Addrs(conf.Consul),
	)

	service := web.NewService(
		web.Name(conf.Service.Name),
		web.Registry(reg),
		web.RegisterTTL(time.Second*REGISTER_TTL),
		web.RegisterInterval(time.Second*REGISTER_INTERVAL),
		web.Address(conf.Service.Addr),
		web.Handler(handler), // use gin's handler
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
