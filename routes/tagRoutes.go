package routes

import (
	"otaviocosta2110/ginClass/controllers"

	"github.com/gin-gonic/gin"
)

func TagRoutes(router *gin.Engine) {
  tagGroup := router.Group("/tag") 
  {
    tagGroup.GET("/getall", controllers.GetAllTags)
  }
}
