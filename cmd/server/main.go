package main

import (
	"log"
	"net/http"

	"tasks-api/internal/handlers"
	customhttp "tasks-api/internal/http"
	"tasks-api/internal/storage"
)

func main() {
	// Инициализация in-memory хранилища
	store := storage.NewMemoryStorage()

	// Инициализация обработчика с хранилищем
	h := handlers.New(store)

	// Настройка маршрутов
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", h.TasksCollection) // GET, POST
	mux.HandleFunc("/tasks/", h.TaskItem)       // GET, PUT, DELETE
	mux.HandleFunc("/health", h.HealthCheck)    // Health check

	// Логирование
	handler := customhttp.LoggingMiddleware(mux)

	log.Println("Сервер запущен на порту :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
