package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
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
func InsertStore() (error, string) {
	db := dbInit()
	defer db.Close()
	ins, err := db.Prepare("INSERT INTO store(store_name,address,open_now,phone_number,website,photo,created_at) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return err, ""
	}
	storeName := "hoge_store_name"
	ins.Exec(storeName, "hoge_address", 1, "hoge_phone_number", "hoge_website", "hoge_photo", time.Now().Format("2006-01-02 03:04:05"))
	return nil, storeName
}

func InsertWiki(storeId int, storeName string) error {
	db := dbInit()
	ins, err := db.Prepare("INSERT INTO wiki(store_id,text,store_user_sum,created_at) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal("insert wiki error")
		return err
	}
	ins.Exec(storeId, storeName, 1, time.Now().Format("2006-01-02 03:04:05"))
	return nil
}

func SelectStore(storeName string) (error, int) {
	db := dbInit()

	var storeId int

	if err := db.QueryRow("SELECT id FROM store WHERE store_name = ?", storeName).Scan(&storeId); err != nil {
		log.Fatal(err)
		return err, storeId
	}

	return nil, storeId
}

func main() {
	err, storeName := InsertStore()
	storeName = "hoge_store_name"
	err, storeId := SelectStore(storeName)
	if err != nil {
		log.Println("エラー")
		log.Println(err)
	}
	fmt.Println(storeId)
}
