package valyria

import (
	"github.com/gin-gonic/gin"
	"github.com/strconv/valyria/config"
	"github.com/strconv/valyria/log"
	"github.com/strconv/valyria/net/server"
)

func RunHTTP(conf *config.Conf, handler *gin.Engine) {
	// HTTP Server
	log.Init(config.C.Service.Log)
	server.InitHTTP(conf, handler)
}
