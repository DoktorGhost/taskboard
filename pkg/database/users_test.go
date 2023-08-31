package database_test

import (
	"taskboard/pkg/database"
	"taskboard/pkg/structs"
	"testing"
)

func TestUserCRUD(t *testing.T) {
	// Создание временной базы данных для тестов
	db := database.InitDB()
	defer db.Close()

	// Инициализация глобальной переменной DB для использования в пакете database
	database.DB = db

	// Тест создания пользователя
	user := &structs.User{
		Username:   "testuser",
		Password:   "testpass",
		First_name: "John",
		Last_name:  "Doe",
	}
	userId, err := database.CreateUser(user)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	// Тест обновления пользователя
	user.First_name = "Jane"
	user.Last_name = "Smith"
	err = database.UpdateUser(userId, user)
	if err != nil {
		t.Errorf("Error updating user: %v", err)
	}

	// Тест удаления пользователя
	err = database.DeleteUser(user.ID)
	if err != nil {
		t.Errorf("Error deleting user: %v", err)
	}
}
