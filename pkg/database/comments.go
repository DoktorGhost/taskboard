package database

import (
	"log"
	"taskboard/pkg/structs"
)

func CreateComment(comment *structs.Comment) (int, error) {
	var commentID int
	err := DB.QueryRow("INSERT INTO comments (description, user_id, task_id) VALUES ($1, $2, $3) RETURNING id",
		comment.Description, comment.UserID, comment.TaskID).Scan(&commentID)
	if err != nil {
		log.Println("Error creating comment:", err)
		return 0, err
	}
	return commentID, nil
}

func UpdateComment(commentID int, comment *structs.Comment) error {
	_, err := DB.Exec("UPDATE comments SET description = $1 WHERE id = $2",
		comment.Description, commentID)
	if err != nil {
		log.Println("Error updating comment:", err)
		return err
	}
	return nil
}

func DeleteComment(commentID int) error {
	_, err := DB.Exec("DELETE FROM comments WHERE id = $1", commentID)
	if err != nil {
		log.Println("Error deleting comment:", err)
		return err
	}
	return nil
}
