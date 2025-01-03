package student_handler

import (
	"net/http"
	"student_management_system/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(ctx *gin.Context){
	id := ctx.Param("id")
	db,exists := ctx.Get("db")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	dbInstance := db.(*gorm.DB)
	var deletedStudent models.Student
	var student models.Student
    if err := dbInstance.First(&student, "id = ?", id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            ctx.JSON(http.StatusNotFound, "No student found with such id")
        } else {
            ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
        }
        return
    }
	res:= dbInstance.Where("id", id).Delete(&deletedStudent)
	 if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	ctx.JSON(http.StatusOK, &student)
}