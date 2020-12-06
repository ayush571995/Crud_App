package main

import (
	"crud_app/internal/database"
	"crud_app/pkg"
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

	studentCrud := pkg.StudentCrud{db}

	http.HandleFunc("/ping", ping)

	http.HandleFunc("/student/insert", studentCrud.InsertToDB)

	http.HandleFunc("/student/delete", studentCrud.DeleteStudent)

	http.ListenAndServe(":8080", nil)

}
