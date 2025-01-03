package main

import (
	"student_management_system/db"
	student_handler "student_management_system/handlers/student"
	"student_management_system/middlewares"
	"student_management_system/models"

	"github.com/gin-gonic/gin"
)

func App(){
	// migrations
	dbInstance:= db.ConnectToDB()
	dbInstance.AutoMigrate(&models.Student{})

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.DbMiddleware())	
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/auth/login")
		apiV1.POST("/auth/register")
		apiV1.GET("/students", student_handler.GetAll)
		apiV1.GET("/student/:id", student_handler.Get)
		apiV1.POST("/student",student_handler.Create)
		apiV1.PATCH("/student/:id")
		apiV1.DELETE("/student/:id", student_handler.Delete)
	}
	router.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, "Running")
	})
	router.Run()

	db.Close(dbInstance)
}