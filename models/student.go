package models

import "github.com/Jeff-Marm-Al/student-registration/db"

type Student struct {
	StudentID int64
	Firstname string `binding:"required"`
	Lastname  string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
}

// struct for the purpose of the GetAllStudents function
type AllStudents struct {
	Firstname string `binding:"required"`
	Lastname  string `binding:"required"`
}

// function to create students
func (new *Student) Save() error {
	// created sql query
	query := "INSERT INTO students(first_name, last_name, email, password) VALUES (?, ?, ?, ?)"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	statementResult, err := statement.Exec(new.Firstname, new.Lastname, new.Email, new.Password)

	if err != nil {
		return err
	}

	studentID, err := statementResult.LastInsertId()

	new.StudentID = studentID

	return err
}

// function to list all students
func GetAllStudents() ([]AllStudents, error) {
	// querying the database for all students
	query := "SELECT * FROM students"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []Student

	for rows.Next() {
		// new student struct for each student
		var student Student
		err := rows.Scan(&student.StudentID, &student.Firstname, &student.Lastname, &student.Email, &student.Password)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	allStudents := make([]AllStudents, len(students))

	for i, student := range students {
		allStudents[i] = AllStudents{
			Firstname: student.Firstname,
			Lastname: student.Lastname,
		}
	}

	return allStudents, nil
}
