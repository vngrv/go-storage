package main

import (
	"fmt"
	"storage"
)

func spawnEmployees(s storage.Storage) {
	for i := 1; i < 10; i++ {
		s.Insert(storage.Employee{Id: i})
	}
}

func main() {
	ms := storage.NewMemoryStorage()

	spawnEmployees(ms)

	fmt.Println(ms.Get(3))

}
