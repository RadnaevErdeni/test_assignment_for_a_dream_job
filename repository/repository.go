package repository

import (
	"tt/testtask"

	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user testtask.Users) (int, error)
	GetAll() ([]testtask.Users, error)
	GetById(userId int) (testtask.Users, error)
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

type Repository struct {
	User
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserDB(db),
		//Task: NewTaskDB(db),
	}
}
