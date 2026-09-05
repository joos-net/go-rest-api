package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"tasks-api/internal/models"
	"tasks-api/internal/storage"
)

// Handler содержит хранилище и методы для обработки HTTP-запросов
type Handler struct {
	Store storage.Storage
}

// New создаем новый экземпляр Handler
func New(s storage.Storage) *Handler {
	return &Handler{Store: s}
}

// writeJSON отправляем ответ в формате JSON
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// writeError отправляем ошибку в JSON формате
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// parseID извлекаем ID из пути
func parseID(w http.ResponseWriter, r *http.Request) (int, bool) {
	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(path)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "Неверный ID задачи")
		return 0, false
	}
	return id, true
}

// TasksCollection обрабатываем запросы к коллекции задач
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTasks(w, r)
	case http.MethodPost:
		h.createTask(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}

// TaskItem обрабатываем запросы по ID
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getTask(w, r, id)
	case http.MethodPut:
		h.updateTask(w, r, id)
	case http.MethodDelete:
		h.deleteTask(w, r, id)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}

// GET /tasks - получение всех задач
func (h *Handler) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.Store.List()
	writeJSON(w, http.StatusOK, tasks)
}

// POST /tasks - создание новой задачи
func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Некорректное тело запроса")
		return
	}
	defer r.Body.Close()

	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		writeError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}

	if task.Title == "" {
		writeError(w, http.StatusBadRequest, "Поле Title обязательно для заполнения")
		return
	}

	createdTask, err := h.Store.Create(task)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Не удалось создать задачу")
		return
	}

	writeJSON(w, http.StatusCreated, createdTask)
}

// GET /tasks/{id} - получение задачи по ID
func (h *Handler) getTask(w http.ResponseWriter, r *http.Request, id int) {
	task, exists := h.Store.Get(id)
	if !exists {
		writeError(w, http.StatusNotFound, fmt.Sprintf("Задача с ID %d не найдена", id))
		return
	}
	writeJSON(w, http.StatusOK, task)
}

// PUT /tasks/{id} - обновление задачи по ID
func (h *Handler) updateTask(w http.ResponseWriter, r *http.Request, id int) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Некорректное тело запроса")
		return
	}
	defer r.Body.Close()

	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		writeError(w, http.StatusBadRequest, "Неверный формат JSON")
		return
	}

	if task.Title == "" {
		writeError(w, http.StatusBadRequest, "Поле Title обязательно для заполнения")
		return
	}

	updatedTask, err := h.Store.Update(id, task)
	if err != nil {
		writeError(w, http.StatusNotFound, fmt.Sprintf("Задача с ID %d не найдена", id))
		return
	}

	writeJSON(w, http.StatusOK, updatedTask)
}

// DELETE /tasks/{id} - удаление задачи по ID
func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	err := h.Store.Delete(id)
	if err != nil {
		writeError(w, http.StatusNotFound, fmt.Sprintf("Задача с ID %d не найдена", id))
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// HealthCheck - проверка состояния сервиса
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}
