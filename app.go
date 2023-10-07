package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("starting RPS service")
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(cors.Default())

	r.POST("/take", handleTake)

	r.Run(":2344")
}
