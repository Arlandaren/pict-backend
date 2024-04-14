package course_endpoints

// In pkg/api/course/list.go

import (
	"course-backend/pkg/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCoursesByTitleHandler fetches courses by title.
func GetCoursesByTitleHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	if title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title query parameter is required"})
		return
	}

	db, exists := ctx.Get("database")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	dbConn, ok := db.(*database.Database)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not of the expected type"})
		return
	}

	courses, err := dbConn.FindCourseByTitle(title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses by title"})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

// GetCoursesByPriceRangeHandler fetches courses within a specified price range.
func GetCoursesByPriceRangeHandler(ctx *gin.Context) {
	minCostStr := ctx.Query("minCost")
	maxCostStr := ctx.Query("maxCost")

	minCost, err := strconv.Atoi(minCostStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minCost query parameter"})
		return
	}

	maxCost, err := strconv.Atoi(maxCostStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maxCost query parameter"})
		return
	}

	db, exists := ctx.Get("database")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	dbConn, ok := db.(*database.Database)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not of the expected type"})
		return
	}

	courses, err := dbConn.FindCoursesByPriceRange(minCost, maxCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses by price range"})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}
