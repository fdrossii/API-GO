package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func Connection() {
	config := mysql.Config{
		User:   "root",
		Passwd: "8841FrPHB!",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "Producto",
	}

	//Establecer la conexión a la BD
	DB, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB CONNECTED")
	}

}
