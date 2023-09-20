package main

type TUserId = int
type TTaskId = int

type Task struct {
	Id          TTaskId `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Completed   bool    `json:"completed"`
}

type TaskStorage map[TUserId]*TasksList

type TasksList struct {
	maxId   TTaskId
	storage map[TTaskId]Task
}
