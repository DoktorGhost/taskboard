package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

// Константы для настройки подключения к базе данных
const (
	DBHost     = "localhost"
	DBPort     = "5432"
	DBUser     = "admin"
	DBPassword = "admin"
	DBName     = "testdb"
)

var DB *sql.DB

// Функция для создания таблиц
func CreateTables() {
	schemaSQL, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	queries := strings.Split(string(schemaSQL), ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query != "" {
			_, err := DB.Exec(query)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// Инициализация базы данных
func InitDB() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)
	var err error
	DB, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
