package main

import (
	"student_management_system/middlewares"

	"github.com/gin-gonic/gin"
)

func App(){
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.DbMiddleware())	
	r.Group("/api/v1/auth")
	{
		r.POST("/login")
		r.POST("/register")
	}
	r.Group("/api/v1/students")
	{
		r.GET("/students")
		r.GET("/student/:id")
		r.POST("/student")
		r.PATCH("/student/:id")
	}
	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, "Running")
	})
	r.Run()
}