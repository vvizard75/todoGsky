package main

import "net/http"

func main() {
	taskRepository := NewTaskRepository()

	httpHandler := NewHttpHandler(taskRepository)

	http.HandleFunc("/users/", httpHandler.ServeHTTP)

	http.ListenAndServe(":8080", nil)
}
