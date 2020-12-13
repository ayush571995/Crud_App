package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crud_app/internal/database"
	"crud_app/internal/graph/generated"
	"fmt"
)

func (r *queryResolver) AllStudents(ctx context.Context) ([]*database.Student, error) {
	var listStudent []*database.Student
	listStudent = append(listStudent, &database.Student{
		FirstName: "Ayush",
		LastName:  "Saluja",
	})

	listStudent = append(listStudent, &database.Student{
		FirstName: "Lol",
		LastName:  "Makeup",
	})

	listStudent = append(listStudent, &database.Student{
		FirstName: "Lol",
		LastName:  "Makeup",
	})

	return listStudent, nil
}

func (r *studentResolver) ID(ctx context.Context, obj *database.Student) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Student returns generated.StudentResolver implementation.
func (r *Resolver) Student() generated.StudentResolver { return &studentResolver{r} }

type queryResolver struct{ *Resolver }
type studentResolver struct{ *Resolver }
