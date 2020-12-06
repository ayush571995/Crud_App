package pkg

import (
	"crud_app/internal/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// StudentCrud is a struct used for holding the crud connection.
type StudentCrud struct {
	Database *sql.DB
}

// InsertToDB is a function used for inserting to the DB.
func (s *StudentCrud) InsertToDB(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}
	log.Println("Request received for inserting the student")

	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var student database.Student
	err = json.Unmarshal(b, &student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	insForm, err := s.Database.Prepare("INSERT INTO student_info(first_name, last_name) VALUES(?,?)")

	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(student.FirstName, student.LastName)

	fmt.Fprintf(w, "Insert Successful")

}

// DeleteStudent is the function for handling the deletion of the student from DB.
func (s *StudentCrud) DeleteStudent(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	id := req.URL.Query()["id"][0]

	log.Println("Request Received for deleting student", id)

	result, err := s.Database.Exec("delete from student_info where id = ?", id)

	if err != nil {

		log.Println("Error while deleting ", err)
		fmt.Fprintf(w, "Error while deletion")

	} else {
		rows, err := result.RowsAffected()
		if err != nil {
			log.Println("Error while deleting ", err)
			fmt.Fprintf(w, "Not able to delete")

		} else if rows == 1 {
			fmt.Fprintf(w, "Deletion Successful")
		} else {
			fmt.Fprintf(w, "No rows affected. Please check Id")
		}

	}
}
