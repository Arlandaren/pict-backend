package router

import (
	"course-backend/pkg/database"
	db_middleware "course-backend/pkg/middleware/db"

	"github.com/gin-gonic/gin"
)

func RouteAPI(r *gin.Engine, db *database.Database) {
	root := r.Group("/api")

	root.Use(db_middleware.HandleBuilder(db))
	{

		couser_methods := root.Group("courses")
		{
			couser_methods.GET("/list/:page")
		}

	}

}
