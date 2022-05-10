package main

import (
	"fmt"
	. "hello.go/awesomeProject/30/pkg"
	"hello.go/awesomeProject/30/pkg/memdb"
	. "hello.go/awesomeProject/30/pkg/storage"
	"log"
	"os"
)

var db Interface

func main() {
	var err error
	pwd := os.Getenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	}
	connstr :=
		"postgres://postgres:" +
			pwd + "@ubuntu-server.northeurope.cloudapp.azure.com/tasks"
	db, err = New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	db = memdb.DB{}
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
}
