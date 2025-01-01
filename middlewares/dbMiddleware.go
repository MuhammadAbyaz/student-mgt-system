package middlewares

import (
	"student_management_system/db"

	"github.com/gin-gonic/gin"
)

func DbMiddleware() gin.HandlerFunc{
	dbInstance := db.ConnectToDB()
	return func(ctx *gin.Context) {
		ctx.Set("db", dbInstance)	
		ctx.Next()
		db.Close(dbInstance)
	}
}