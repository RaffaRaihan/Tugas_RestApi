package usercontroller

import (
	"encoding/json"
	"net/http"
	"tugas_restApi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ValidateToken(r *http.Request) bool {
    token := r.Header.Get("Authorization")
    return token == "Bearer token_app"

}

func Index(c *gin.Context) {
	var user []models.Users

	models.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func Show(c *gin.Context) {
	var user []models.Users
	id := c.Param("id")

	if err := models.DB.First(&user, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"massage": "Id item kosong"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"massage": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}

func Create(c *gin.Context) {

	var user models.Users

	if err := c.ShouldBindJSON(&user); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return
	}
	
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Update(c *gin.Context) {

	var user models.Users
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return
	}

	if models.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"massage": "Tidak dapat di Update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"massage": "Data berhasil di Update"})


}

func Delete(c *gin.Context) {

	var user models.Users

	var input struct{
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&user, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"massage": "Tidak dapat di hapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"massage": "Berhasil di hapus"})
}