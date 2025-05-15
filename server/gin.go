// gin_server.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func startGinServer() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		msg := c.Query("msg")
		if msg == "" {
			msg = "Nothing received (did you forget ?msg=... ?)"
		}
		fmt.Println("Received:", msg)
		c.String(200, "Hello Client, I got: %s", msg)
	})

	r.Run(":8080") // 启动 Gin 服务
}
