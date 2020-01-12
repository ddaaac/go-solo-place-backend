package main

import (
	"./Config"
	"./Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Config.ConnectDB()
	defer Config.DisconnectDB()
	Routes.SetupRouter(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
