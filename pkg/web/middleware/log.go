package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		if raw != "" {
			path = path + "?" + raw
		}

		httpInfo := fmt.Sprintf("[GIN] | %3d | %13v | %15s | %-7s %#v",
			c.Writer.Status(),
			end.Sub(start).String(),
			c.ClientIP(),
			c.Request.Method,
			path,
		)

		logrus.Infof("[Gin]: %s", httpInfo)
	}
}
