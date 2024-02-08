package studentcontroller

import (
	"net/http"

	"github.com/bryanagamk/crud-go-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var students []models.Student

	models.DB.Where("deleted_at IS NULL").Find(&students)
	c.JSON(http.StatusOK, gin.H{"data": students})
}

func Store(c *gin.Context) {}

func Show(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if err := models.DB.First(&student, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

func Update(c *gin.Context) {}
func Delete(c *gin.Context) {}
