package storage

import "hello.go/awesomeProject/30/pkg/storage"

// Интерфейс БД

type Interface interface {
	Tasks(int, int) ([]storage.Task, error)
	NewTask(storage.Task) (int, error)
}
