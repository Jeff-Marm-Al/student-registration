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
	query := "SELECT first_name, last_name FROM students"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allStudents []AllStudents

	for rows.Next() {
		// new student struct for each student
		var student AllStudents
		err := rows.Scan(&student.Firstname, &student.Lastname)

		if err != nil {
			return nil, err
		}

		allStudents = append(allStudents, student)
	}

	return allStudents, nil
}

func GetStudentInfo(firstname, lastname string) (*Student, error) {
	// querying database for all students where the firstname and lastname match what is provided
	query := "SELECT * FROM students WHERE first_name = ? AND last_name = ?"
	row:= db.DB.QueryRow(query, firstname, lastname)

	var student Student

	err := row.Scan(&student.StudentID, &student.Firstname, &student.Lastname, &student.Email, &student.Password)

	if err != nil {
		return nil, err
	}

	return &student, nil
}
