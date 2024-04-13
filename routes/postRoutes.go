package routes

import (
	"otaviocosta2110/ginClass/controllers"

	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.Engine) {
  postGroup := router.Group("/post") 
  {
    postGroup.POST("/create", controllers.CreatePost)
    postGroup.GET("/getbyclass/:class", controllers.GetPostByClass)
  }
}

