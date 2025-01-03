package student_handler

import (
	"net/http"
	"student_management_system/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(ctx *gin.Context){
	id:= ctx.Param("id")
	db, exists:= ctx.Get("db")
	if !exists{
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	dbInstance := db.(*gorm.DB)
	var student models.Student
	response := dbInstance.Where("id = ?", id).First(&student)
	if response.Error == gorm.ErrRecordNotFound{
		ctx.JSON(http.StatusNotFound, "No student found with such id")
		return
	} else if response.Error != nil{
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return	
	}
	ctx.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}