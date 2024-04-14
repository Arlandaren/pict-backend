package utils

import (
	"course-backend/pkg/database"
	"fmt"
	"math/rand"
	"time"
)

func convertStepikCourseToCourse(stepikCourse StepikCourse_) database.Course {

	// Map fields from StepikCourse_ to Course
	// Note: You might need to adjust the mapping based on the actual data and how you want to represent it in your Course struct

	priceList := []int{0, 0, 0, 0, 0, 990, 100, 3999, 4999, 1000, 120, 1100, 1200}
	rateList := []float64{0.4, 1, 2.9, 2, 3, 3.5, 4.5, 4.8, 4.9, 5}

	return database.Course{
		ID:          stepikCourse.Id,
		Title:       stepikCourse.Title,
		Description: stepikCourse.Summary,                 // Assuming the summary is the description
		Link:        stepikCourse.Cover,                   // Assuming the cover is the link
		Complexity:  rand.Intn(5) + 1,                     // Assuming the review summary is the complexity
		Cost:        priceList[rand.Intn(len(priceList))], // Assuming the learners count is the cost
		Rating:      rateList[rand.Intn(len(rateList))],
	}
}

func RunWatcher(db *database.Database) {
	for {
		newCourses := getStepikCourses()
		fmt.Println("Sync courses")
		// Add new courses to the database
		for _, course := range newCourses {
			if err := db.AddCourse(convertStepikCourseToCourse(course)); err != nil {
				// Handle error
				// For example, log the error
			}
		}

		// Sleep for an hour
		time.Sleep(time.Hour)
	}
}
