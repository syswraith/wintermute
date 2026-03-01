package minerva

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect()(*sql.DB) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/wintermute")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB!")
	return db
}

func Ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB Pinged!")
}

func SelectAll(db *sql.DB) {

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

		fmt.Println(shorturl, longurl)
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
	}

}

