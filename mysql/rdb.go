package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@tcp(go_tutorial_mysql)/test")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("insert into todo values(?,?,?,?)")
	if err != nil{
		log.Fatal(err)
	}
	_, err = stmt.Exec("2", "test", "6/8", "6/8")
	if err != nil{
		log.Fatal(err)
	}

	var id int
	var task string
	rows, err := db.Query("select id, task from todo")
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){
		err := rows.Scan(&id, &task)
		if err != nil{
			log.Fatal(err)
		}
		log.Println(id, task)
	}
	err = rows.Err()
	if err != nil{
		log.Fatal(err)
	}

}
