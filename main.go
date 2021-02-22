package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("cluster-tool/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"xff":       c.Request.Header.Get("X-Forwarded-For"),
			"client IP": c.ClientIP(),
		})
	})

	r.Run()
}
