package routes

import (
	"otaviocosta2110/ginClass/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
  userGroup := router.Group("/user") 
  {
    userGroup.GET("/getall", controllers.GetAllUsers)
    userGroup.GET("/getbyemail/:email", controllers.GetUserByEmail)
    userGroup.POST("/create", controllers.CreateUser)
  }
}
