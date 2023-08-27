package database

import (
	"log"
	"taskboard/pkg/structs"
)

// Функция для создания нового пользователя
func CreateUser(user *structs.User) error {
	_, err := DB.Exec("INSERT INTO users (username, password, first_name, last_name) VALUES ($1, $2, $3, $4)",
		user.Username, user.Password, user.First_name, user.Last_name)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	return nil
}

// Функция для обновления информации о пользователе
func UpdateUser(user *structs.User) error {
	_, err := DB.Exec("UPDATE users SET username = $1, password = $2, first_name = $3, last_name = $4 WHERE id = $5",
		user.Username, user.Password, user.First_name, user.Last_name, user.ID)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}
	return nil
}

// Функция для удаления пользователя
func DeleteUser(userID int) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	return nil
}
