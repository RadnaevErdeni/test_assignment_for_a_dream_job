package service

import (
	"tt/repository"
	"tt/testtask"
)

type User interface {
	Create(user testtask.DBUsers) (int, error)
	GetAll(surname, name, patronymic, address string, id, passportSerie, passportNumber, limit, offset int) ([]testtask.DBUsers, error)
	GetById(userId int) (testtask.DBUsers, error)
	Delete(userId int) error
	Update(userId int, input testtask.UpdateUserInput) error
	LaborCosts(userId int, start, end *string) ([]testtask.LaborCosts, error)
}

type Task interface {
	Create(userId int, task testtask.Tasks) (int, error)
	GetAll(userId int) ([]testtask.Tasks, error)
	GetById(userId, taskId int) (testtask.Tasks, error)
	Delete(userId, taskId int) error
	Start(userId, taskId int) error
	End(userId, taskId int) error
	Pause(userId, taskId int) error
	Resume(userId, taskId int) error
}

type Service struct {
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Task: NewTaskService(repos.Task, repos.User),
	}
}
