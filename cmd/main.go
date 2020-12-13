package main

import (
	"crud_app/internal/database"
	"crud_app/internal/graph"
	"crud_app/internal/graph/generated"
	"crud_app/pkg"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.ListenAndServe(":8080", nil)

}
