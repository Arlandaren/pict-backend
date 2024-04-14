package db_middleware

import (
	"course-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

func HandleBuilder(db *database.Database) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.Set("database", db)
		ctx.Next()
	}
}
