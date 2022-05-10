package memdb

import "hello.go/awesomeProject/30/pkg/storage"

type DB []storage.Task

func (db DB) Tasks(int, int) ([]storage.Task, error) {
	return db, nil
}
func (db DB) NewTask(storage.Task) (int, error) {
	return 0, nil
}
