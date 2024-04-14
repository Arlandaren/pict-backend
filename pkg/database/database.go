package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) FindCourseByName(name string) ([]Course, error) {
	rows, err := d.db.Query("SELECT * FROM courses WHERE title LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.Link, &course.Complexity, &course.Cost, &course.Rating); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}

// In pkg/database/database.go

// FindCourseByTitle finds courses by title.
func (d *Database) FindCourseByTitle(title string) ([]Course, error) {
	rows, err := d.db.Query("SELECT * FROM courses WHERE title LIKE ?", "%"+title+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.Link, &course.Complexity, &course.Cost, &course.Rating); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}

// FindCoursesByPriceRange finds courses within a specified price range.
func (d *Database) FindCoursesByPriceRange(minCost, maxCost int) ([]Course, error) {
	rows, err := d.db.Query("SELECT * FROM courses WHERE cost >= ? AND cost <= ?", minCost, maxCost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.Link, &course.Complexity, &course.Cost, &course.Rating); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}

func (d *Database) AddCourse(course Course) error {
	// Assuming the ID is auto-incremented by the database
	_, err := d.db.Exec("INSERT INTO courses (title, description, link, complexity, cost, rating) VALUES (?, ?, ?, ?, ?, ?)",
		course.Title, course.Description, course.Link, course.Complexity, course.Cost, course.Rating)
	return err
}

type Course struct {
	ID          int
	Title       string
	Description string
	Link        string
	Complexity  int
	Cost        int
	Rating      float64
}
