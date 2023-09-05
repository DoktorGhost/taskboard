package main

import (
	"fmt"
	"net/http"
	"taskboard/pkg/api"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	// Установите маршрут для главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "СЕРВЕР ЗАПУЩЕН 2")
	})

	// Установите маршрут для обработчика RegisterHandler
	http.HandleFunc("/register", api.RegisterHandler)

	// Установите обработчик для статических файлов
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Запустите сервер
	server.ListenAndServe()
}
