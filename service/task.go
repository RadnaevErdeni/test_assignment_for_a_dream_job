package service

import (
	"tt/repository"
)

type TaskService struct {
	repo     repository.User
	listRepo repository.Task
}

func NewTaskService(repo repository.User, listRepo repository.Task) *TaskService {
	return &TaskService{repo: repo, listRepo: listRepo}
}
