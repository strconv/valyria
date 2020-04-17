package server

import (
	"github.com/gin-gonic/gin"
)

func NewHTTP() *gin.Engine {
	return initHandler()
}

func initHandler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router.Use() //middleware trace、swagger、
	return router
}
