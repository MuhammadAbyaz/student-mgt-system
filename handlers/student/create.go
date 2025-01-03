package student_handler

import (
	"encoding/json"
	"net/http"
	"student_management_system/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func Create(ctx *gin.Context){
	db, exists:= ctx.Get("db")
	if !exists{
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	dbInstance := db.(*gorm.DB)
	var newStudent models.Student
	err := json.NewDecoder(ctx.Request.Body).Decode(&newStudent)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	res := dbInstance.Create(&newStudent)
	if res.Error != nil {
		if validationErr := res.Error.Error(); validationErr != "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": validationErr,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	ctx.JSON(http.StatusCreated, &newStudent)
}