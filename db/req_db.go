package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func dbInit() *sql.DB {
	godotenv.Load()
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_TCP")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal("db error.")
	}
	return db
}

//TODO lat,lng 追加
func InsertStore() error {
	db := dbInit()
	defer db.Close()
	ins, err := db.Prepare("INSERT INTO store(store_name,address,open_now,phone_number,website,photo) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	ins.Exec("hoge_store_name", "hoge_address", 1, "hoge_phone_number", "hoge_website", "hoge_photo")
	return nil
}

func InsertWiki() error {
	db := dbInit()
	ins, err := db.Prepare("INSERT INTO store(hoge,hoge,hoge,hoge) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	ins.Exec("golang-2019", "golang+001@gmail.com", "Jhon", "123456")
	return nil
}

func SelectStore() {
	db := dbInit()

	storeName := "hoge_store_name"

	var result int

	if err := db.QueryRow("SELECT id FROM store WHERE store_name = ?", storeName).Scan(&result); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func main() {
	//err := InsertStore()
	SelectStore()
	//if err != nil {
	//	log.Println("エラー")
	//	log.Println(err)
	//}
}
