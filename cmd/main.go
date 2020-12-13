package main

import (
	"crud_app/internal/database"
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {

	println("Starting server")

	db, err := database.InitDB()

	if err != nil {
		panic("DB connection not successful")
	}

	studentCrud := database.NewUserDBImpl(db)

	studentFetcher := database.InitStudentCrud(studentCrud)

	http.HandleFunc("/ping", ping)

	http.HandleFunc("/student/insert", studentFetcher.InsertToDB)

	http.HandleFunc("/student/getByID", studentFetcher.GetByID)

	http.HandleFunc("/student/getAll", studentFetcher.GetAll)

	http.ListenAndServe(":8080", nil)

}
