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

func (s *UserService) Create(user testtask.DBUsers) (int, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetById(id int) (testtask.DBUsers, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}
func (s *UserService) Update(userId int, input testtask.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, input)
}

func (s *UserService) GetAll(surname, name, patronymic, address string, id, passportSerie, passportNumber, limit, offset int) ([]testtask.DBUsers, error) {
	return s.repo.GetAll(surname, name, patronymic, address, id, passportSerie, passportNumber, limit, offset)
}
func (s *UserService) LaborCosts(userId int, start, end *string) ([]testtask.LaborCosts, error) {
	return s.repo.LaborCosts(userId, start, end)
}
