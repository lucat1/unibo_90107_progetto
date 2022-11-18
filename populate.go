package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=db user=postgres password=postgres dbname=tommaso sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	schema, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		log.Fatalf("Could not read schema.sql: %v", err)
	}

	res, err := db.Exec("DROP SCHEMA public CASCADE;" + "CREATE SCHEMA public;")
	if err != nil {
		log.Fatalf("Could not purge all data: %v", err)
	}
	res, err = db.Exec(string(schema))
	if err != nil {
		log.Fatalf("Could not apply schema.sql: %v", err)
	}
	fmt.Println(res)
}
