package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("mysql", "root:software2022@tcp(localhost:3306)/ecommerce_go")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error al verificar la conexión:", err)
	}

	log.Println("Conexión a la base de datos exitosa.")
}

func GetDB() *sql.DB {
	return DB
}
