package main

import (
	"net/http"
	"otaviocosta2110/ginClass/routes"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context){
  c.IndentedJSON(http.StatusOK, gin.H{"message": "tudo show"})
}

func main(){
  router := gin.Default()

  router.GET("/healthcheck", healthCheck)
  routes.UserRoutes(router)
  router.Run(":8080")
}
