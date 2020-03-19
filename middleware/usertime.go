// creator: zacyuan
// date: 2019-12-29
// 耗时日志中间简

package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuanzhangcai/microgin/log"
)

// UsedTime 生成耗时日志中间件
func UsedTime() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		body, _ := c.GetRawData()

		resp, _ := c.Get("response")

		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency = param.Latency - param.Latency%time.Second
		}
		log.Info(fmt.Sprintf("UsedTime: %3d| %13v |%s %-7s %s body[%s] resp[%v] err[%s]",
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			string(body),
			resp,
			param.ErrorMessage,
		))
	}
}
