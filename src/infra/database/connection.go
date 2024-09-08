package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var SQL *sql.DB

// Função para inicializar a conexão com o banco
func InitDB() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	SQL = db
}
