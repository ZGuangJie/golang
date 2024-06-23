package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Loader struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

func loginVerify(c *gin.Context) {
	var load Loader
	c.ShouldBind(&load)
	fmt.Printf("%#v", load)
	if load.Name == "zhu" && load.Password == "123456" {
		c.Set("example", "Recived")
		c.Next()
	} else {
		c.Abort()
		c.Redirect(http.StatusMovedPermanently, "/error")
	}
}

func main() {
	r := gin.New()
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	r.Use(gin.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	r.Static("/static", "./static")
	// 解析HTML模板
	r.LoadHTMLGlob("./template/*.html")
	// 相当于访问 /v1/login, ect
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	login := r.Group("/index")
	login.Use(loginVerify)
	{
		login.POST("/submit", func(c *gin.Context) {
			message, ok := c.Get("example")
			if !ok {
				message = "Null"
			}
			c.JSON(http.StatusOK, gin.H{"URI": message})
		})
		r.GET("/error", func(c *gin.Context) {
			c.String(http.StatusOK, "用户名或者密码错误")
		})
	}

	err := r.Run(":9090")
	if err != nil {
		log.Fatal("Service run fail...")
	}
}
