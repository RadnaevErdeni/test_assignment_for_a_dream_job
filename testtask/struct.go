package testtask

import (
	"errors"
	"strings"
)

type Users struct {
	Id              int    `json:"id" db:"id"`
	Passport_number string `json:"passport_number" db:"passport_number" binding:"required"`
	Surname         string `json:"surname" db:"surname"`
	Name            string `json:"name" db:"name" binding:"required"`
	Patronymic      string `json:"patronymic" db:"patronymic"`
	Address         string `json:"address" db:"address"`
}

type Tasks struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Start_time  string `json:"start_time" db:"start_time" binding:"required"`
	End_time    string `json:"end_time" db:"end_time" binding:"required"`
	Duration    string `json:"duration" db:"duration"`
}

type UserTask struct {
	Id      int
	User_id int
	Task_id int
}

type UpdateTaskInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Start_time  string  `json:"start_time" db:"start_time"`
	End_time    string  `json:"end_time" db:"end_time"`
	Duration    string  `json:"duration" db:"duration"`
}

func (i Users) ValidatePasNum() error {
	pn := i.Passport_number
	if len(pn) != 11 {
		return errors.New("invalid passport number")
	}
	sn := strings.Split(pn, " ")
	if len(sn) != 2 {
		return errors.New("invalid passport number")
	}
	if len(sn[0]) != 4 && len(sn[1]) != 6 {
		return errors.New("invalid passport number")
	}
	return nil
}
func (i UpdateTaskInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateUserInput struct {
	Passport_number *string `json:"passport_number"`
	Surname         *string `json:"surname" `
	Name            *string `json:"name" `
	Patronymic      *string `json:"patronymic"`
	Address         *string `json:"address"`
}

func (i UpdateUserInput) Validate() error {
	if i.Passport_number == nil && i.Surname == nil && i.Name == nil && i.Patronymic == nil && i.Address == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
