package database

import (
	"taskboard/pkg/structs"
	"testing"
	"time"
)

func TestCommentCRUD(t *testing.T) {
	db := InitDB()
	defer db.Close()

	DB = db

	// Создаем тестового пользователя
	user := &structs.User{
		Username:   "COMMENTATORS",
		Password:   "testpass",
		First_name: "Комментатор",
		Last_name:  "НОМЕРОДИН",
	}
	userId, err := CreateUser(user)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	// Создаем тестовую задачу
	task := &structs.Task{
		Title:       "Тестовая задача комментов",
		Description: "В этой задаче нужно написать комментарии. Действуй!",
		Priority:    structs.PriorityLow,
		Status:      structs.StatusPending,
		DueDate:     time.Now(),
	}

	idTask, err := CreateTask(task)
	if err != nil {
		t.Errorf("Error creating task: %v", err)
	}

	// Создаем тестовый комментарий
	comment := &structs.Comment{
		Description: "Это тестовый комменатрий, нужно изменить что-то в задаче",
		UserID:      userId,
		TaskID:      idTask,
		DueDate:     time.Now(),
	}

	commentID, err := CreateComment(comment)
	if err != nil {
		t.Errorf("Error creating comment: %v", err)
	}

	// Тестируем обновление комментария
	comment.Description = "Это комменатрий обновлен, но его удалят!"
	err = UpdateComment(commentID, comment)
	if err != nil {
		t.Errorf("Error updating comment: %v", err)
	}

	comment.Description = "Это комменатрий не должен быть удален"
	_, err = CreateComment(comment)
	if err != nil {
		t.Errorf("Error creating comment: %v", err)
	}

	// Тестируем удаление комментария
	err = DeleteComment(commentID)
	if err != nil {
		t.Errorf("Error deleting comment: %v", err)
	}
}
