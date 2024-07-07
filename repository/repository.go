package repository

import (
	"tt/testtask"

	"github.com/jmoiron/sqlx"
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
	UpdateTask(userId, taskId int, input testtask.UpdateTaskInput) error
	Start(userId, taskId int) error
	End(userId, taskId int) error
}

type Repository struct {
	User
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserDB(db),
		Task: NewTaskDB(db),
	}
}
