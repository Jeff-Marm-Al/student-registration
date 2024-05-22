package models

import "github.com/Jeff-Marm-Al/student-registration/db"

type Course struct {
	CourseID           int64
	Title              string
	Ticker             string
	InstructorLastName string
	Credits            int64
	Capacity           int64
}

// function to insert a new course to the database
func (new *Course) Save() error {
	// created sql query
	query := "INSERT INTO courses(title, ticker, instructor_last_name, capacity) VALUES (?, ?, ?, ?)"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	statementResult, err := statement.Exec(new.Title, new.Ticker, new.InstructorLastName, new.Capacity)

	if err != nil {
		return err
	}

	classID, err := statementResult.LastInsertId()

	new.CourseID = classID

	return err
}

/*
func GetAllCourses() {

}
*/
