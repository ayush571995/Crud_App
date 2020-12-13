package database

// StudentFetcher interface is interface providing CRUD functionalities
type StudentFetcher interface {
	InsertToDB(student Student) error

	GetAllStudents() ([]Student, error)

	FetchByID(id int64) (*Student, error)
}
