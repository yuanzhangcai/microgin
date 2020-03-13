// creator: zacyuan
// date: 2019-12-30

package middleware

import (
	"github.com/gin-gonic/gin"
)

// Accessor 添加跨域访问控制
func Accessor() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		c.Next()
	}
}
