// references
// https://github.com/diegothucao/rest-api-golang
// https://github.com/chefgs/golang/tree/master/src/api

package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/command/status", GetCommandStatus)
	router.GET("/api/command", PostCommands)
	router.PUT("/api/command/{taskId}", PostCancelCommands)

	listenPort := os.Getenv("appPort")
	if listenPort == "" {
		listenPort = "8081"
	}
	router.Run(":" + listenPort)
}