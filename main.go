package main

import (
	studentcontroller "github.com/bryanagamk/crud-go-api/controllers/studentController"
	"github.com/bryanagamk/crud-go-api/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))

	r.GET("/students", studentcontroller.Index)
	r.POST("/students", studentcontroller.Store)
	r.GET("/students/:id", studentcontroller.Show)
	r.PATCH("/students/:id", studentcontroller.Update)
	r.DELETE("/students/:id", studentcontroller.Delete)

	r.Run()
}
