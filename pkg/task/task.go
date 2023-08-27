package task

import (
	"errors"
	ts "taskboard/pkg/structs"
)

var tasks []ts.Task
var lastID int

// AddTask добавляет новую задачу
func AddTask(task ts.Task) {
	task.ID = lastID
	lastID++
	tasks = append(tasks, task)
}

// UpdateTask обновляет задачу по ID
func UpdateTask(id int, updatedTask ts.Task) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			return nil
		}
	}
	return errors.New("task not found")
}

// DeleteTask удаляет задачу по ID
func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

// GetTasks возвращает список всех задач
func GetTasks() []ts.Task {
	return tasks
}
