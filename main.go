package main

import (
	"fmt"
	"lark/pkg/common/xgin"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := xgin.NewGinServer()
	engine.Engine.GET("hellow", hello)
	engine.Run(9166)
}

func hello(c *gin.Context) {
	fmt.Println("访问了hello这个api")
	c.SecureJSON(0, "访问成功")
}
