package service

import (
	"tt/repository"
	"tt/testtask"
)

type TaskService struct {
	repo     repository.Task
	listRepo repository.User
}

func NewTaskService(repo repository.Task, listRepo repository.User) *TaskService {
	return &TaskService{repo: repo, listRepo: listRepo}
}

func (s *TaskService) Create(userId int, task testtask.Tasks) (int, error) {
	return s.repo.Create(userId, task)
}
func (s *TaskService) GetById(userId, taskId int) (testtask.Tasks, error) {
	return s.repo.GetById(userId, taskId)
}

func (s *TaskService) Delete(userId, taskId int) error {
	return s.repo.Delete(userId, taskId)
}
func (s *TaskService) UpdateTask(userId, taskId int, input testtask.UpdateTaskInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateTask(userId, taskId, input)
}

func (s *TaskService) GetAll(userId int) ([]testtask.Tasks, error) {
	return s.repo.GetAll(userId)
}
func (s *TaskService) Start(userId, taskId int) error {
	return s.repo.Start(userId, taskId)
}
func (s *TaskService) End(userId, taskId int) error {
	return s.repo.End(userId, taskId)
}
func (s *TaskService) Pause(userId, taskId int) error {
	return s.repo.Pause(userId, taskId)
}

func (s *TaskService) Resume(userId, taskId int) error {
	return s.repo.Resume(userId, taskId)
}
