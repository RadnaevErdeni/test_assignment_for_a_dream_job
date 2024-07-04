package repository

import (
	"fmt"
	"strings"
	"tt/testtask"

	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{db: db}
}
func (r *UserDB) Create(user testtask.Users) (int, error) {
	var id int
	createuser := fmt.Sprintf("INSERT INTO %s (passport_number, surname,name,patronymic,address) values ($1, $2,$3,$4,$5) RETURNING id", usersTable)
	row := r.db.QueryRow(createuser, user.Passport_number, user.Surname, user.Name, user.Patronymic, user.Address)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *UserDB) GetAll() ([]testtask.Users, error) {
	var users []testtask.Users
	query := fmt.Sprintf(`SELECT us.id, us.passport_number,us.surname,us.name,us.patronymic,us.address FROM %s us`,
		usersTable)
	err := r.db.Select(&users, query)
	return users, err
}

func (r *UserDB) GetById(id int) (testtask.Users, error) {
	var user testtask.Users
	query := fmt.Sprintf(`SELECT us.id, us.passport_number,us.surname,us.name,us.patronymic,us.address FROM %s us WHERE us.id = $1`,
		usersTable)
	err := r.db.Get(&user, query, id)
	return user, err
}
func (r *UserDB) Delete(userId int) error {
	deleteStr := fmt.Sprintf("DELETE FROM %s WHERE id =$1", usersTable)
	_, err := r.db.Exec(deleteStr, userId)
	if err != nil {
		return err
	}
	return err
}
func (r *UserDB) Update(userId int, input testtask.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Passport_number != nil {
		setValues = append(setValues, fmt.Sprintf("passport_number=$%d", argId))
		args = append(args, *input.Passport_number)
		argId++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}
	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Patronymic != nil {
		setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
		args = append(args, *input.Patronymic)
		argId++
	}
	if input.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, *input.Address)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s us SET %s WHERE us.id = $%d",
		usersTable, setQuery, argId)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
