package main

import (
	"fmt"
	"log"

	"github.com/upper/db/v4/adapter/postgresql"
)

var settings = postgresql.ConnectionURL{

	Database: "dnerd",
	Host:     "localhost:5433",
	User:     "dnerd",
	Password: "",
}

func main() {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("postgresql.Open: ", err)
	}
	defer sess.Close()

	// The `db` API is portable, you can expect code to work the same on
	// different databases
	howManyBooks, err := sess.Collection("books").Find().Count()
	if err != nil {
		log.Fatal("Find: ", err)
	}

	fmt.Printf("We have %d books in our database.\n", howManyBooks)
}
