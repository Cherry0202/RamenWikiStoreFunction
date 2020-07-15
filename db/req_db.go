package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DbInit() *sql.DB {
	godotenv.Load()
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_TCP")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal("db error.")
	}

	defer db.Close()

	return db
}

func InsertStore() {
	db := DbInit()
	ins, err := db.Prepare("INSERT INTO store(hoge,hoge,hoge,hoge) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec("golang-2019", "golang+001@gmail.com", "Jhon", "123456")
	return
}

func InsertWiki() {
	db := DbInit()
	ins, err := db.Prepare("INSERT INTO store(hoge,hoge,hoge,hoge) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec("golang-2019", "golang+001@gmail.com", "Jhon", "123456")
	return
}
