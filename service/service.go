package service

import (
	"tt/repository"
	"tt/testtask"
)

type User interface {
	Create(user testtask.Users) (int, error)
	GetAll() ([]testtask.Users, error)
	GetById(id int) (testtask.Users, error)
	Delete(userId int) error
	Update(userId int, input testtask.UpdateUserInput) error
}

type Task interface {
	/*Create(userId, taskId int, item testtask.Task) (int, error)
	GetAll(userId, taskId int) ([]testtask.Task, error)
	GetById(userId, taskId int) (testtask.Task, error)
	Delete(userId, taskId int) error
	UpdateTasks(userId, taskId int, input testtask.UpdateTaskInput) error*/
}

type Service struct {
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		/*Task: NewTaskService(repos.Task, repos.User),*/
	}
}
