package database

import (
	"taskboard/pkg/structs"
	"testing"
	"time"
)

func TestTaskCRUD(t *testing.T) {
	// Создание временной базы данных для тестов
	db := InitDB()
	defer db.Close()

	DB = db

	// Тест создания задачи
	task := &structs.Task{
		Title:       "Test Task",
		Description: "Description test task",
		DueDate:     time.Now(),
		Priority:    structs.PriorityLow,
		Status:      structs.StatusPending,
	}
	id, err := CreateTask(task)
	if err != nil {
		t.Errorf("Error creating task: %v", err)
	}

	// Тест обновления задачи
	task.Title = "Updated Task"
	task.Description = "Updated Description task"
	err = UpdateTask(id, task)
	if err != nil {
		t.Errorf("Error updating task: %v", err)
	}

	// Тест удаления задачи
	err = DeleteTask(task.ID)
	if err != nil {
		t.Errorf("Error deleting task: %v", err)
	}
}
