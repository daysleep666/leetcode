package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dev")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	t := time.Now()
	tx, _ := db.Begin()
	for i := 0; i < 100000; i++ {
		a := fmt.Sprintf("%d", i)
		t = t.Add(-time.Minute)
		_, err := tx.Exec("insert into book (name, author, create_at) values (?,?,?)", a, a, t)

		// INSERT INTO book (name, author, create_at, is_precious) VALUES ('小明奇遇记', '小明', '2018-01-01', '1');

		if err != nil {
			panic(err)
		}
	}
	tx.Commit()
}
