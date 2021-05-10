package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func RunApp() {

	routes()

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
