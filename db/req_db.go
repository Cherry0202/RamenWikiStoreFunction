package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func dbInit() *sql.DB {
	godotenv.Load()
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_TCP")+":"+os.Getenv("DB_CONNECTION_PORT")+")/"+os.Getenv("DB_NAME"))

	if err != nil {
		log.Println(err, "in db init error")
	}
	return db
}

func InsertStore(storeName string, storeAddress string, openNow int, phoneNumber string, webSite string, photoRef string, lat float64, lng float64, openTime string) (error, string) {
	db := dbInit()
	defer db.Close()
	ins, err := db.Prepare("INSERT INTO store(store_name,address,open_now,phone_number,website,photo,lat,lng,open_time,created_at) VALUES(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err, "in insert store error")
		return err, ""
	}
	ins.Exec(storeName, storeAddress, openNow, phoneNumber, webSite, photoRef, lat, lng, openTime, time.Now().Format("2006-01-02 03:04:05"))
	return nil, storeName
}

func InsertWiki(storeId int, storeName string) error {
	db := dbInit()
	ins, err := db.Prepare("INSERT INTO wiki(store_id,text,store_user_sum,created_at) VALUES(?,?,?,?)")
	if err != nil {
		log.Println(err, "in insert wiki error")
		return err
	}
	ins.Exec(storeId, storeName, 1, time.Now().Format("2006-01-02 03:04:05"))
	return nil
}

func SelectStore(storeName string) (error, int) {
	db := dbInit()

	var storeId int

	if err := db.QueryRow("SELECT id FROM store WHERE store_name = ?", storeName).Scan(&storeId); err != nil {
		log.Println(err, "in select store error")
		return err, storeId
	}

	return nil, storeId
}
