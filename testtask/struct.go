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
	Patronymic      string `json:"patronymic" db:"patronymic"`
	Address         string `json:"address" db:"address" binding:"required"`
}

type Users struct {
	Id              int    `json:"id"`
	Passport_number string `json:"passport_number" binding:"required"`
	Surname         string `json:"surname" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Patronymic      string `json:"patronymic" `
	Address         string `json:"address" binding:"required"`
}
type Times struct {
	Start_time *string `json:"start_time"`
	End_time   *string `json:"end_time"`
}

type LaborCosts struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic" `
	Title      string `json:"title"`
	Duration   string `json:"duration"`
}
type Tasks struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Start_time  string `json:"start_time" db:"start_time"`
	End_time    string `json:"end_time" db:"end_time"`
	Duration    string `json:"duration" db:"duration"`
	Done        bool   `json:"done" db:"done"`
	Took        bool   `json:"took" db:"took"`
	Date_create string `json:"date_create" db:"date_create"`
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

/*
	func (i *UpdateUserInput) ValidatePasNumUp(usr UpdateUserInput) error {
		if len(*usr.Passport_number) != 6 && len(*usr.Passport_serie) != 4 {
			return errors.New("invalid passport number")
		}
		return nil
	}
*/
func (i *Users) ValidatePasNum(usr Users) (DBUsers, error) {
	var user DBUsers
	if len(usr.Passport_number) != 11 {
		return user, errors.New("invalid passport number")
	}
	sn := strings.Split(usr.Passport_number, " ")
	if len(sn) != 2 {
		return user, errors.New("invalid passport number")
	}
	if len(sn[0]) != 4 && len(sn[1]) != 6 {
		return user, errors.New("invalid passport number")
	}
	pn, err := strconv.Atoi(sn[0])
	if err != nil {
		return user, err
	}
	ps, err := strconv.Atoi(sn[1])
	if err != nil {
		return user, err
	}
	user.Surname = usr.Surname
	user.Name = usr.Name
	user.Patronymic = usr.Patronymic
	user.Address = usr.Address
	user.Passport_number = ps
	user.Passport_serie = pn
	return user, nil
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
