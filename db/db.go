package db

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
)

// global database variable so it can be used outside of the package
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Could not connect to database" + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	createTables()
}

func createTables() {
	createStudentsTable := `
	CREATE TABLE IF NOT EXISTS students (
		student_id INTEGER(4) NOT NULL PRIMARY KEY AUTOINCREMENT,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password VARCHAR(20) NOT NULL
	)
	`

	_, err := DB.Exec(createStudentsTable)

	if err != nil {
		panic("Could not create students table: " + err.Error())
	}

	createCoursesTable := `
	CREATE TABLE IF NOT EXISTS courses (
		course_id INTEGER(4) NOT NULL PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(50) NOT NULL,
		ticker VARCHAR(10) NOT NULL,
		instructor_last_name VARCHAR(15) NOT NULL,
		credits INTEGER NOT NULL,
		capacity INTEGER NOT NULL 
	)
	`

	_, err = DB.Exec(createCoursesTable)

	if err != nil {
		panic("Could not create courses table: " + err.Error())
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		registration_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		student_id INTEGER(4),
		course_id INTEGER(4),
		registration_date DATE,
		FOREIGN KEY (student_id) REFERENCES students(student_id),
		FOREIGN KEY (course_id) REFERENCES courses(course_id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table: " + err.Error())
	}
}