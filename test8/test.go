package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	sql := "DELETE FROM \"test\" WHERE age = "
	sprintf := fmt.Sprintf("%s%d", sql, 1)
	fmt.Println(sprintf)

	exec, err2 := db.Exec(sprintf)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("effect:", exec)
}

var (
	err error
	db  *sql.DB
)

const (
	pgHost     = "b2904236d6.zicp.vip"
	pgPort     = 22021
	pgUser     = "root"
	pgPassword = "@ps"
	// pgDbname   = "blockdb"
	pgDbname = "testdb"
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPassword, pgDbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Connect PG Failed: ", err)
	}

	db.SetMaxOpenConns(50)

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping GP Failed: ", err)
	}
}
