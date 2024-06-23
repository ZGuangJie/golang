package route

import (
	"fmt"
	"lastPrograma/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = gin.New()

func RounteInit() {
	// 手动添加 Logger 和 Recovery 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	loadFiles()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func loadFiles() {
	r.Static("/static", "./static")
	r.LoadHTMLFiles("./templates/index.html")
}
func RounteStart() {
	// 根路由，返回主界面
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	v1 := r.Group("/v1")
	{
		// 增加 post
		v1.POST("/todo", func(ctx *gin.Context) {
			// 1. 前端填写待办事项，点击提交，发送到这里
			var item model.Todo
			err := ctx.ShouldBind(&item)
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("%#v", item)
			// 2. 拿数据，并存入数据库
			// 3. 返回响应
		})
		// 删除 delete
		v1.DELETE("/todo/:id", func(ctx *gin.Context) {

		})
		// 修改 put
		v1.PUT("/todo/:id", func(ctx *gin.Context) {

		})
		// 查看指定Id的item get
		v1.GET("/todo/:id", func(ctx *gin.Context) {

		})
		// 查看所有的items get
		v1.GET("/todo", func(ctx *gin.Context) {
			var items []model.Todo
			items = append(items, model.Todo{
				Id:     1,
				Title:  "吃饭",
				Status: false,
			})
			ctx.JSON(http.StatusOK, items)
		})
	}
}
