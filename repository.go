package main

import (
	"errors"
	"sync"
)

type IDataRepository interface {
	GetAll(userId int) ([]Task, error)
	Create(userId int, task *Task) error
	Update(userId int, task *Task) error
	Delete(userId int, id int) error
}

type TaskRepository struct {
	mainStorage TaskStorage
	mu          sync.RWMutex
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		mainStorage: make(TaskStorage),
	}
}

func (t TaskRepository) GetAll(userId int) ([]Task, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	userStorage, ok := t.mainStorage[userId]
	if !ok {
		return nil, errors.New("user not found")
	}

	var tasks []Task
	for _, task := range userStorage.storage {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t TaskRepository) Create(userId int, task *Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	userStorage, ok := t.mainStorage[userId]
	if !ok {
		userStorage = &TasksList{
			maxId:   0,
			storage: make(map[TTaskId]Task),
		}
		t.mainStorage[userId] = userStorage
	}

	userStorage.maxId++
	task.Id = userStorage.maxId
	userStorage.storage[userStorage.maxId] = Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		Completed:   task.Completed,
	}
	return nil
}

func (t TaskRepository) Update(userId int, task *Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	userStorage, ok := t.mainStorage[userId]
	if !ok {
		return errors.New("user not found")
	}

	_, ok = userStorage.storage[task.Id]
	if !ok {
		return errors.New("task not found")
	}

	userStorage.storage[task.Id] = Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		Completed:   task.Completed,
	}
	return nil
}

func (t TaskRepository) Delete(userId int, id int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	userStorage, ok := t.mainStorage[userId]
	if !ok {
		return errors.New("user not found")
	}

	_, ok = userStorage.storage[id]
	if !ok {
		return errors.New("task not found")
	}

	delete(userStorage.storage, id)
	return nil
}
