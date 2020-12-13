package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// StudentCrud is a struct used for holding the crud connection.
type StudentCrud struct {
	Fetcher StudentFetcher
}

// InitStudentCrud gives you the Db function.
func InitStudentCrud(fetcherImpl StudentFetcher) StudentCrud {
	return StudentCrud{Fetcher: fetcherImpl}
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
	var student Student
	err = json.Unmarshal(b, &student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = s.Fetcher.InsertToDB(student)

	if err != nil {
		http.Error(w, "Insert Failed", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Insert Successful")

}

// GetByID is the function for handling the get by id of the student from DB.
func (s *StudentCrud) GetByID(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}

	id := req.URL.Query()["id"][0]

	log.Println("Request Received for get by id student", id)

	decimalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
	}
	result, err := s.Fetcher.FetchByID(decimalID)

	if err != nil {
		http.Error(w, "Get BY Id failed", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

//GetAll fetches all the users registered in DB.
func (s *StudentCrud) GetAll(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid_http_method")
		return
	}
	result, err := s.Fetcher.GetAllStudents()
	if err != nil {
		http.Error(w, "Get All failed", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
