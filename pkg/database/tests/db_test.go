package database

import (
	"database/sql"
	"taskboard/pkg/database"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateTables(t *testing.T) {
	// Создание временной базы данных для тестов
	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=testdb sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	database.DB = db

	database.CreateTables()

	// Проверка, что таблицы были созданы
	expectedTables := []string{"users", "roles", "user_roles", "tasks", "comments"}

	for _, tableName := range expectedTables {
		rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_name = $1", tableName)
		if err != nil {
			t.Errorf("Error querying table %s: %v", tableName, err)
		}
		defer rows.Close()

		if !rows.Next() {
			t.Errorf("Table %s was not created", tableName)
		}
	}
}

func TestInitDB(t *testing.T) {
	// Тестируем инициализацию базы данных
	db := database.InitDB()

	// Проверяем, что база данных успешно инициализировалась
	if db == nil {
		t.Error("Database initialization failed")
	}
}
