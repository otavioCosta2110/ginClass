package routes

import (
	"otaviocosta2110/ginClass/controllers"

	"github.com/gin-gonic/gin"
)

func ClassRoutes(router *gin.Engine){
  classGroup := router.Group("/class")
  {
    classGroup.GET("/getall", controllers.GetAllClasses)
    classGroup.POST("/create", controllers.CreateClass)
  }
}
