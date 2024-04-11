package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context){
  c.IndentedJSON(http.StatusOK, gin.H{"message": "tudo show"})
}

func main(){
  router := gin.Default()

  router.GET("/healthcheck", healthCheck)
  router.Run(":8080")
}
