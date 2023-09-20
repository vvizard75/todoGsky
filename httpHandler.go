package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type HttpHandler struct {
	repository IDataRepository
}

func NewHttpHandler(repository IDataRepository) *HttpHandler {
	return &HttpHandler{
		repository: repository,
	}
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.getAllTasks(w, r)
	case http.MethodPost:
		h.createTask(w, r)
	case http.MethodPut:
		h.updateTask(w, r)
	case http.MethodDelete:
		h.deleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (h HttpHandler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[1] != "users" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks, err := h.repository.GetAll(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonString, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonString)
}

func (h HttpHandler) createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[1] != "users" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var task Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.repository.Create(userId, &task)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(strconv.Itoa(task.Id)))
}

func (h HttpHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 || parts[1] != "users" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	taskId, err := strconv.Atoi(parts[3])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var task Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task.Id = taskId
	err = h.repository.Update(userId, &task)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h HttpHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 || parts[1] != "users" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	taskId, err := strconv.Atoi(parts[3])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.repository.Delete(userId, taskId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
