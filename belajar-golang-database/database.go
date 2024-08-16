package belajar_golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10) // minimal koneksi
	db.SetMaxOpenConns(100) // maksimal koneksi
	db.SetConnMaxIdleTime(5 * time.Minute) // berapa lama koneksi yang akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) // berapa lama koneksi yang akan digunakan

	return db
}
