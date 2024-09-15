package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/css", "./css")
	router.Static("/images", "./css/images")

	router.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	err := router.Run(":2435")
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
