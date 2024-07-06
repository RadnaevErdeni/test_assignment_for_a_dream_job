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
func (r *UserDB) Create(user testtask.DBUsers) (int, error) {
	var id int
	createuser := fmt.Sprintf("INSERT INTO %s (passport_serie, passport_number,  surname,name,patronymic,address) values ($1, $2,$3,$4,$5,$6) RETURNING id", usersTable)
	row := r.db.QueryRow(createuser, user.Passport_serie, user.Passport_number, user.Surname, user.Name, user.Patronymic, user.Address)
	if err := row.Scan(&id); err != nil {
		return id, nil
	}

	return id, nil
}

func (r *UserDB) GetAll(surname, name, patronymic, address string, id, passportSerie, passportNumber, limit, offset int) ([]testtask.DBUsers, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if id != 0 {
		setValues = append(setValues, fmt.Sprintf("id=$%d", argId))
		args = append(args, id)
		argId++
	}
	if surname != "" {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, surname)
		argId++
	}
	if name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, name)
		argId++
	}
	if patronymic != "" {
		setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
		args = append(args, patronymic)
		argId++
	}
	if passportSerie != 0 {
		setValues = append(setValues, fmt.Sprintf("passport_serie=$%d", argId))
		args = append(args, passportSerie)
		argId++
	}
	if passportNumber != 0 {
		setValues = append(setValues, fmt.Sprintf("passport_number=$%d", argId))
		args = append(args, passportNumber)
		argId++
	}
	if address != "" {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, address)
		argId++
	}

	whereClause := ""
	if len(setValues) > 0 {
		whereClause = "WHERE " + strings.Join(setValues, " AND ")
	}

	query := fmt.Sprintf("SELECT id, surname, name,patronymic, passport_serie, passport_number,address FROM users %s LIMIT $%d OFFSET $%d",
		whereClause, argId, argId+1)
	args = append(args, limit, offset)

	var users []testtask.DBUsers
	err := r.db.Select(&users, query, args...)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserDB) GetById(id int) (testtask.DBUsers, error) {
	var user testtask.DBUsers
	query := fmt.Sprintf(`SELECT us.id,us.passport_serie, us.passport_number,us.surname,us.name,us.patronymic,us.address FROM %s us WHERE us.id = $1`,
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

	if input.Passport_serie != nil {
		setValues = append(setValues, fmt.Sprintf("passport_serie=$%d", argId))
		args = append(args, *input.Passport_serie)
		argId++
	}

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

func (r *UserDB) LaborCosts(userId int, start, end string) ([]testtask.LaborCosts, error) {
	var lc []testtask.LaborCosts
	query := fmt.Sprintf(`SELECT surname,name,patronymic,title,TO_CHAR(justify_hours(duration), 'HH24:MI') as duration FROM %s ut INNER JOIN %s us ON us.id = ut.user_id INNER JOIN %s  ts ON ts.id = ut.task_id WHERE us.id = $1 AND start_time >= $2 AND end_time <= $3 ORDER BY duration DESC`,
		usersTaskTable, usersTable, taskTable)
	if err := r.db.Select(&lc, query, userId, start, end); err != nil {
		fmt.Println(err)
		fmt.Println(start)
		return nil, err
	}
	return lc, nil
}

/*
func (r *UserDB) LaborCosts(userId int, start, end string) ([]testtask.LaborCosts, error) {
	var lc []testtask.LaborCosts
	query := `
		SELECT surname, name, patronymic, title, TO_CHAR(justify_hours(duration), 'HH24:MI') as duration
		FROM users_tasks ut
		INNER JOIN users us ON us.id = ut.user_id
		INNER JOIN tasks ts ON ts.id = ut.task_id
		WHERE us.id = $1 AND start_time >= $2 AND end_time <= $3
		ORDER BY duration DESC`

	// Преобразуем строки в тип time.Time
	startTime, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return nil, fmt.Errorf("invalid start time format: %v", err)
	}
	endTime, err := time.Parse(time.RFC3339, end)
	if err != nil {
		return nil, fmt.Errorf("invalid end time format: %v", err)
	}

	// Выполнение запроса с параметрами
	if err := r.db.Select(&lc, query, userId, startTime, endTime); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return lc, nil
}
*/
