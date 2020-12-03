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

type StudentCrud struct {
	Database *sql.DB
}

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
