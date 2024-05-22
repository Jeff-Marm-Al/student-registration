package models

import "github.com/Jeff-Marm-Al/student-registration/db"

type Course struct {
	CourseID           int64
	Title              string `binding:"required"`
	Ticker             string `binding:"required"`
	InstructorLastName string `binding:"required"`
	Credits            int64  `binding:"required"`
	Capacity           int64  `binding:"required"`
}

type AllCourses struct {
	Ticker  string
	Title   string
	Credits int64
}

// function to insert a new course to the database
func (new *Course) Save() error {
	// created sql query
	query := "INSERT INTO courses(title, ticker, instructor_last_name, credits, capacity) VALUES (?, ?, ?, ?, ?)"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	statementResult, err := statement.Exec(new.Title, new.Ticker, new.InstructorLastName, new.Credits, new.Capacity)

	if err != nil {
		return err
	}

	classID, err := statementResult.LastInsertId()

	new.CourseID = classID

	return err
}

// function to query for all courses available
func GetAllCourses() ([]AllCourses, error) {
	query := "SELECT ticker, title, credits FROM courses"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var courses []AllCourses

	for rows.Next() {
		var course AllCourses

		err := rows.Scan(&course.Ticker, &course.Title, &course.Credits)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)

	}

	return courses, nil
}
