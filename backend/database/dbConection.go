package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func GetConectionDB() {
	HOST := os.Getenv("HOST")
	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")

	dsn := USER + ":" + PASSWORD + "@tcp(" + HOST + ")/" + DBNAME
	var err error
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error comprobando la base de datos:", err)
	}

	log.Println("Conexión a la base de datos exitosa")

}
