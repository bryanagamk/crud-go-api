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

func Store(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&student)
	c.JSON(http.StatusOK, gin.H{"student": student})
}

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

func Update(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if err := c.ShouldBindJSON(&student); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	idAffected := models.DB.Model(&student).Where("id_student = ?", id).Updates(&student).RowsAffected

	if idAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak dapat mengupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"student": "Data berhasil diubah"})
}

func Delete(c *gin.Context) {}
