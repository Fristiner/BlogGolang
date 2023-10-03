package router

import (
	"Blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/assets", "./assets")

	//r.GET("/index", controller.ListUser)
	r.POST("/register", controller.RegisterUser)
	r.GET("/register", controller.GoRegister)
	r.GET("/", controller.Index)
	r.GET("/login", controller.GoLogin)
	r.POST("/login", controller.Login)
	// 博客界面
	r.GET("/post_index", controller.GetPostIndex)
	r.POST("/post", controller.AddPost)
	r.GET("/post", controller.GoAddPost)

	r.GET("/detail", controller.PostDetail)
	_ = r.Run()
}
