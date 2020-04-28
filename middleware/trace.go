package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/strconv/valyria/trace"
)

// trace
func TracerWrapper(c *gin.Context) {
	trace.TracerWrapper(c)
}
