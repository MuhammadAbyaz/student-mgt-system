package middlewares

import (
	"student_management_system/db"

	"github.com/gin-gonic/gin"
)

func DbMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		dbInstance := db.ConnectToDB()
		ctx.Set("db", dbInstance)	
		ctx.Next()
	}
}