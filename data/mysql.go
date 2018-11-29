package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "gouser:gopass@/twitter?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT tweet_id, content FROM tweets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tweet_id int
		var content string

		if err := rows.Scan(&tweet_id, &content); err != nil {
			log.Fatal(err)
		}
		fmt.Println(tweet_id, content)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
