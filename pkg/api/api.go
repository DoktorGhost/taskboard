package api

import (
	"encoding/json"
	"net/http"
	"taskboard/pkg/database"
	"taskboard/pkg/structs"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		// Обработка ошибки
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Вызов функции CreateUser для вставки пользователя в базу данных
	userID, err := database.CreateUser(&user)
	if err != nil {
		// Если произошла ошибка при вставке в базу данных, верните соответствующий HTTP-статус и сообщение об ошибке
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": "Ошибка при сохранении пользователя в базе данных"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Теперь у вас есть ID нового пользователя (userID) после успешной регистрации
	// Отправьте ответ в JSON формате с ID пользователя
	response := map[string]int{"user_id": userID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
