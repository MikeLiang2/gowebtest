// beego_server.go
package main

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	msg := c.GetString("msg")
	if msg == "" {
		msg = "Nothing received (did you forget ?msg=... ?)"
	}
	fmt.Println("Received:", msg)
	c.Ctx.WriteString("Hello Client, I got: " + msg)
}

func startBeegoServer() {
	web.Router("/", &MainController{})
	web.Run()
}
