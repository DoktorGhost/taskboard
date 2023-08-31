package database

import (
	"log"
	"taskboard/pkg/structs"
)

// Функция для создания новой задачи
func CreateTask(task *structs.Task) (int, error) {
	var taskID int
	err := DB.QueryRow("INSERT INTO tasks (title, description, priority, due_date, status) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		task.Title, task.Description, task.Priority, task.DueDate, task.Status).Scan(&taskID)
	if err != nil {
		log.Println("Error creating task:", err)
		return 0, err
	}
	return taskID, nil
}

// Функция для обновления информации о задаче
func UpdateTask(taskID int, task *structs.Task) error {
	_, err := DB.Exec("UPDATE tasks SET title = $1, description = $2, priority = $3, due_date = $4, status = $5 WHERE id = $6",
		task.Title, task.Description, task.Priority, task.DueDate, task.Status, taskID)
	if err != nil {
		log.Println("Error updating task:", err)
		return err
	}
	return nil
}

// Функция для удаления задачи
func DeleteTask(taskID int) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id = $1", taskID)
	if err != nil {
		log.Println("Error deleting task:", err)
		return err
	}
	return nil
}
