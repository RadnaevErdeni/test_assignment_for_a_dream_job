package testtask

import (
	"errors"
	"strconv"
	"strings"
)

type DBUsers struct {
	Id              int    `json:"id" db:"id"`
	Passport_serie  int    `json:"passport_serie" db:"passport_serie" binding:"required"`
	Passport_number int    `json:"passport_number" db:"passport_number" binding:"required"`
	Surname         string `json:"surname" db:"surname"`
	Name            string `json:"name" db:"name" binding:"required"`
	Patronymic      string `json:"patronymic,omitempty" db:"patronymic"`
	Address         string `json:"address" db:"address" binding:"required"`
}
type Passport struct {
	PassportNumber string `json:"passport_number"`
}

type Times struct {
	Start_time *string `json:"start_time"`
	End_time   *string `json:"end_time"`
}

type LaborCosts struct {
	Surname     string  `json:"surname"`
	Name        string  `json:"name"`
	Patronymic  string  `json:"patronymic" `
	Title       string  `json:"title"`
	Duration    *string `json:"duration,omitempty"`
	Count_pause *int    `json:"count_pause,omitempty"`
}
type Tasks struct {
	Id             int     `json:"id" db:"id"`
	Title          string  `json:"title,omitempty" db:"title" binding:"required"`
	Description    string  `json:"description,omitempty" db:"description"`
	Start_time     *string `json:"start_time,omitempty" db:"start_time"`
	End_time       *string `json:"end_time,omitempty" db:"end_time"`
	Duration       *string `json:"duration,omitempty" db:"duration"`
	Duration_pause *string `json:"duration_pause,omitempty" db:"duration_pause"`
	Done           bool    `json:"done,omitempty" db:"done"`
	Took           bool    `json:"took,omitempty" db:"took"`
	Date_create    string  `json:"date_create" db:"date_create"`
	Pause_time     *string `json:"pause_time,omitempty" db:"pause_time"`
	Resume_time    *string `json:"resume_time,omitempty" db:"resume_time"`
	Count_pause    int     `json:"count_pause" db:"count_pause"`
}

type UserTask struct {
	Id      int
	User_id int
	Task_id int
}
type EndTask struct {
	End_time *string `json:"end_time" db:"end_time"`
	Done     *bool   `json:"bool" db:"bool"`
}
type UpdateTaskInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Start_time  *string `json:"start_time"`
	End_time    *string `json:"end_time"`
	Duration    *string `json:"duration"`
	Done        *bool   `json:"bool"`
	Took        *bool   `json:"took"`
}

func (i *Passport) ValidatePasNum(usr Passport) (int, int, error) {
	if len(usr.PassportNumber) != 11 {
		return 0, 0, errors.New("invalid passport number")
	}
	sn := strings.Split(usr.PassportNumber, " ")
	if len(sn) != 2 {
		return 0, 0, errors.New("invalid passport number")
	}
	if len(sn[0]) != 4 && len(sn[1]) != 6 {
		return 0, 0, errors.New("invalid passport number")
	}
	pn, err := strconv.Atoi(sn[1])
	if err != nil {
		return 0, 0, err
	}
	ps, err := strconv.Atoi(sn[0])
	if err != nil {
		return 0, 0, err
	}

	return ps, pn, nil
}

func (i UpdateTaskInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Start_time == nil && i.End_time == nil && i.Duration == nil && i.Done == nil && i.Took == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateUserInput struct {
	Passport_serie  *string `json:"passport_serie"`
	Passport_number *string `json:"passport_number"`
	Surname         *string `json:"surname" `
	Name            *string `json:"name" `
	Patronymic      *string `json:"patronymic"`
	Address         *string `json:"address"`
}

func (i UpdateUserInput) Validate() error {
	if len(*i.Passport_number) != 6 {
		return errors.New("invalid passport number")
	}
	if len(*i.Passport_serie) != 4 {
		return errors.New("invalid passport number")
	}
	if i.Passport_serie == nil && i.Passport_number == nil && i.Surname == nil && i.Name == nil && i.Patronymic == nil && i.Address == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
