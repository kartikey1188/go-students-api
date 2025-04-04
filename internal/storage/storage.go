package storage

import "github.com/kartikey1188/go-students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudent(id int64) (types.Student, error)
}
