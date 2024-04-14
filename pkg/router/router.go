package router

import (
	course_endpoints "course-backend/pkg/api/course"
	"course-backend/pkg/database"
	db_middleware "course-backend/pkg/middleware/db"

	"github.com/gin-gonic/gin"
)

func RouteAPI(r *gin.Engine, db *database.Database) {
	root := r.Group("/api")

	root.Use(db_middleware.HandleBuilder(db))
	{
		courseMethods := root.Group("courses")
		{
			// Route for getting courses by title
			courseMethods.GET("/by-title", course_endpoints.GetCoursesByTitleHandler)

			// Route for getting courses within a price range
			courseMethods.GET("/by-price-range", course_endpoints.GetCoursesByPriceRangeHandler)
		}
	}
}
