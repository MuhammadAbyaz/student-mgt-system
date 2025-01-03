package student_handler

import (
	"net/http"
	"student_management_system/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(ctx *gin.Context){
	db, exists := ctx.Get("db")
	if !exists {	
		ctx.JSON(http.StatusInternalServerError, "No Database Instance Found")
		return
	}
	dbInstance := db.(*gorm.DB)
	var students []models.Student
	res:= dbInstance.Find(&students)
	if res.Error != nil{
		ctx.JSON(http.StatusInternalServerError, res.Error.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}