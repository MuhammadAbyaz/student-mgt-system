package main

import "github.com/gin-gonic/gin"

func App(){
	router := gin.Default()
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, "Running")
	})
	router.Run()
}