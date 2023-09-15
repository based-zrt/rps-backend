package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("starting RPS service")
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.POST("/take", handleTake)

	r.Run(":2344")
}
