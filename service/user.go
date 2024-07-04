package service

import (
	"tt/repository"
	"tt/testtask"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user testtask.Users) (int, error) {
	if err := user.ValidatePasNum(); err != nil {
		return 0, err
	}
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]testtask.Users, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(id int) (testtask.Users, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}
func (s *UserService) Update(userId int, input testtask.UpdateUserInput) error {
	return s.repo.Update(userId, input)
}
