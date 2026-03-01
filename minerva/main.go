package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/wintermute")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db connected")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pinged")



	var (
		shorturl string
		longurl string
	)

	rows, err := db.Query("select * from links;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&shorturl, &longurl)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(shorturl, longurl)
	}


	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}


	defer db.Close()
}
