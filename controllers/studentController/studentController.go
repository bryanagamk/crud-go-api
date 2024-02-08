package studentcontroller

import (
	"net/http"

	"github.com/bryanagamk/crud-go-api/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var students []models.Student

	models.DB.Find(&students)
	c.JSON(http.StatusOK, gin.H{"students": students})
}

func Store(c *gin.Context)  {}
func Show(c *gin.Context)   {}
func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
