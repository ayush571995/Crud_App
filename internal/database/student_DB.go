package database

import (
	"database/sql"
	"fmt"
	"log"
)

//StudentDB struct is the impl of the fetcher interface.
type StudentDB struct {
	Database *sql.DB
}

// NewUserDBImpl gives the Underlying infra for DB.
func NewUserDBImpl(database *sql.DB) StudentDB {
	return StudentDB{Database: database}
}

// InsertToDB function inserts the student record in DB.
func (s StudentDB) InsertToDB(student Student) error {
	insForm, err := s.Database.Prepare("INSERT INTO student_info(first_name, last_name) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = insForm.Exec(student.FirstName, student.LastName)
	if err != nil {
		log.Println("Unable to insert getting error ", err)
		return err
	}
	return nil
}

//GetAllStudents function fetches all the students from DB.
func (s StudentDB) GetAllStudents() ([]Student, error) {
	rows, err := s.Database.Query("select * from student_info")
	if err != nil {
		log.Println("Unable to get all students ", err)
		return nil, err
	}
	defer rows.Close()
	var pk int32
	var firstName string
	var lastName string
	var studentList []Student
	for rows.Next() {
		err := rows.Scan(&pk, &firstName, &lastName)
		if err != nil {
			log.Println("Error while scanning")
			return nil, err
		}
		studentList = append(studentList, Student{FirstName: firstName, LastName: lastName})
	}
	return studentList, nil
}

//FetchByID function fetches the specific record from the DB.
func (s StudentDB) FetchByID(id int64) (*Student, error) {
	rows, err := s.Database.Query("select * from student_info where id = ?", id)
	if err != nil {
		log.Println("Unable to get student ", id, err)
		return nil, err
	}
	defer rows.Close()
	var pk int32
	var firstName string
	var lastName string
	for rows.Next() {
		err := rows.Scan(&pk, &firstName, &lastName)
		if err != nil {
			log.Println("Error while scanning")
			return nil, err
		}
		return &Student{FirstName: firstName, LastName: lastName}, nil
	}
	log.Println("No matching Id found for id ", id)
	return nil, fmt.Errorf("Invalid Id ")
}
