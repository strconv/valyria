package middleware

import (
	"net/http"

	"github.com/strconv/valyria/jwt"

	"github.com/gin-gonic/gin"
)

// 跨域
func CORS() gin.HandlerFunc {
	// https://blog.csdn.net/u010918487/article/details/82686293
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,"+jwt.HEADER_TOKEN_KEY)
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
