package main

import (
	studentcontroller "github.com/bryanagamk/crud-go-api/controllers/studentController"
	"github.com/bryanagamk/crud-go-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/students", studentcontroller.Index)
	r.POST("/students", studentcontroller.Store)
	r.GET("/students/:id", studentcontroller.Show)
	r.PATCH("/students/:id", studentcontroller.Update)
	r.DELETE("/students/:id", studentcontroller.Delete)

	r.Run()
}
