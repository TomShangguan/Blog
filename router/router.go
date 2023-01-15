package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/post_index", controller.Index)

	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.GET("/", controller.GetPostIndex)
	e.GET("/post", controller.GoAddPost)
	e.POST("/post", controller.AddPost)

	e.GET("/detail", controller.PostDetail)

	e.Run()
}
