package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type students struct {
	Name string `form:"name"`
	Age  int16  `form:"age"`
}

func main() {

	r := gin.Default()
	// r.Static("/static", "./static")

	r.LoadHTMLGlob("template/*")
	// r.LoadHTMLFiles("template/form.html")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		// fmt.Printf("%v", name)
		c.JSON(http.StatusOK, c.Request.URL)
	})
	r.POST("/index", func(c *gin.Context) {
		var stu students
		err := c.ShouldBind(&stu)
		if err != nil {
			c.JSON(http.StatusBadGateway, "Params delivery ERROR!")
		} else {
			fmt.Printf("%#v\n", stu)
			c.JSON(http.StatusOK, stu)
		}
	})

	r.GET("/upload", func(c *gin.Context) {
		// fmt.Printf("%v", name)
		c.HTML(http.StatusOK, "upload.html", nil)
	})
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := filepath.Join("./uploadFiles", file.Filename)
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.GET("/test/:name/*action", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.URL.Path)
	})

	r.Run(":9090")
}
