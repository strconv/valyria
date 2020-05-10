package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/strconv/valyria/config"
	"github.com/strconv/valyria/jwt"
	"github.com/strconv/valyria/log"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := log.For(c, "middleware", "jwt_auth")

		token := c.Request.Header.Get(jwt.HEADER_TOKEN_KEY)
		if token == "" {
			respInvalidJWT(c, "missing jwt token")
			c.Abort() // 中间件内的返回须abort，否则会继续进行后面的请求
			return    // 因为return的作用域在本层
		}
		logger.Info("get token: ", token)
		claims, err := jwt.ParseToken(token)
		if err != nil {
			respInvalidJWT(c, err.Error())
			c.Abort()
			return
		}

		// 续签（需要端上对每个响应判断是否含有 HEADER_TOKEN_KEY 对应字段，有的话更新jwt_token）
		if time.Now().Unix()-claims.ExpiresAt < int64(config.C.JWT.Timeout/20) {
			token, _ = jwt.GenToken(claims.UID)
			c.Header(jwt.HEADER_TOKEN_KEY, token)
		}

		c.Set("uid", claims.UID) // 后续的请求使用 c.GET("uid")  获取uid
		c.Next()
	}
}

func respInvalidJWT(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"msg":  "JWT鉴权错误: " + msg,
	})
}
