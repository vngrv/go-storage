package storage

import (
	"errors"
)

type Employee struct {
	Id     int
	Name   string
	Age    string
	Salary int
}

type Storage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error
}

type MemoryStorage struct {
	data map[int]Employee
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *MemoryStorage) Insert(e Employee) error {
	s.data[e.Id] = e
	return nil
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	e, exist := s.data[id]

	if !exist {
		return Employee{}, errors.New("employee with such id dosn't exist")
	}

	return e, nil
}

func (s *MemoryStorage) Delete(id int) error {
	delete(s.data, id)
	return nil
}
