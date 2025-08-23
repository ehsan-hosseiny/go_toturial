package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "user=admin dbname=test sslmode=disable port=5454 password=123456 ")
	if err != nil {
		log.Fatal(err)
	}

	

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)
	if err != nil {
		log.Fatal(err)
	}

	// res, err := db.Exec(`INSERT INTO users (username,email) VALUES($1,$2)`, "ehsan2", "ehsanhossini2@gmail.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	rows , err := db.Query(`SELECT id,username,email,created_at FROM users;`)

	defer db.Close()

	for rows.Next(){
		var id int
		var username,email string
		var createAd string
		err := rows.Scan(&id,&username,&email,&createAd)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Printf("id = %d , username = %s, email = %s createAd = %s\n",id,username,email,createAd)
	}

	
	

}
