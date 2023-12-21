package main

import (
	"github.com/gin-gonic/gin"
	"hello/controllers"
)

func main() {
	engine := gin.Default()

	engine.GET("/message", controllers.GetMessage)
	engine.POST("/message", controllers.AddMessage)
	engine.POST("/update", controllers.ChangeMessage)
	engine.DELETE("/message/:id", controllers.Deletemessage)

	engine.Run("127.0.0.1:8080")
}
