package storage

import (
	"fmt"
	"refactor28/pkg/student"
)

type Storage map[string]*student.Students

func New() Storage {
	return make(map[string]*student.Students)
}

func (storage Storage) Put(student *student.Students) {
	storage[student.Name] = student
}

func (storage Storage) Get(studentName string) (*student.Students, error) {
	stud, ok := storage[studentName]
	if !ok {
		return nil, fmt.Errorf("no such user")
	}
	return stud, nil
}
