package repository

import (
	"fmt"
	"strings"
	"tt/testtask"

	"github.com/jmoiron/sqlx"
)

type TaskDB struct {
	db *sqlx.DB
}

func NewTaskDB(db *sqlx.DB) *TaskDB {
	return &TaskDB{db: db}
}

func (r *TaskDB) Create(userId int, task testtask.Tasks) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description,done,took,date_create,duration) values ($1, $2, false,false,Now(),'00:00:00') RETURNING id", taskTable)
	row := tx.QueryRow(createTaskQuery, task.Title, task.Description)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserTaskQuery := fmt.Sprintf("INSERT INTO %s (user_id, task_id) values ($1, $2)", usersTaskTable)
	_, err = tx.Exec(createUserTaskQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TaskDB) GetAll(userId int) ([]testtask.Tasks, error) {
	var tasks []testtask.Tasks
	query := fmt.Sprintf("SELECT ts.id,ts.title,ts.description,ts.start_time,ts.end_time,ts.duration,ts.done,ts.took,ts.date_create FROM %s ts INNER JOIN %s ut ON ts.id = ut.task_id WHERE ut.user_id = $1", taskTable, usersTaskTable)
	if err := r.db.Select(&tasks, query, userId); err != nil {
		return nil, err
	}
	return tasks, nil
}
func (r *TaskDB) GetById(userId, taskId int) (testtask.Tasks, error) {
	var user testtask.Tasks
	query := fmt.Sprintf(`SELECT ts.id,ts.title,ts.description,ts.start_time,ts.end_time,ts.duration,ts.done,ts.took,ts.date_create FROM %s ts INNER JOIN user_task us ON us.task_id = ts.id WHERE ts.id = $1 AND us.user_id = $2`,
		taskTable)
	err := r.db.Get(&user, query, taskId, userId)
	return user, err
}
func (r *TaskDB) Delete(userId, taskId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	deleteUserTask := fmt.Sprintf("DELETE FROM %s us WHERE us.user_id = $1 AND us.task_id =$2", usersTaskTable)
	_, err = tx.Exec(deleteUserTask, userId, taskId)
	if err != nil {
		tx.Rollback()
		return err
	}
	deleteTask := fmt.Sprintf("DELETE FROM %s WHERE id =$1 ", taskTable)
	_, err = tx.Exec(deleteTask, taskId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TaskDB) UpdateTask(userId, taskId int, input testtask.UpdateTaskInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Start_time != nil {
		setValues = append(setValues, fmt.Sprintf("start_time=$%d", argId))
		args = append(args, *input.Start_time)
		argId++
	}
	if input.End_time != nil {
		setValues = append(setValues, fmt.Sprintf("end_time=$%d", argId))
		args = append(args, *input.End_time)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	if input.Took != nil {
		setValues = append(setValues, fmt.Sprintf("took=$%d", argId))
		args = append(args, *input.Took)
		argId++
	}
	if input.Duration != nil {
		setValues = append(setValues, fmt.Sprintf("duration=$%d", argId))
		args = append(args, *input.Duration)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ts SET %s WHERE ts.id = $%d",
		taskTable, setQuery, argId)
	args = append(args, taskId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TaskDB) Start(userId, taskId int) error {
	startQuery := fmt.Sprintf("UPDATE %s ts SET start_time = Now(),took = true FROM %s tsk,%s ut WHERE ts.id = $1 AND ut.task_id = ts.id AND ut.user_id = $2", taskTable, taskTable, usersTaskTable)
	_, err := r.db.Exec(startQuery, taskId, userId)
	if err != nil {
		return err
	}

	return nil
}
func (r *TaskDB) End(userId, taskId int) error {
	endQuery := fmt.Sprintf("UPDATE %s ts SET end_time = Now(),done = true ,duration = Now() - ts.start_time FROM %s tsk, %s ut WHERE ts.id = $1 AND ut.task_id = ts.id AND ut.user_id = $2", taskTable, taskTable, usersTaskTable)
	_, err := r.db.Exec(endQuery, taskId, userId)
	if err != nil {
		return err
	}

	return nil
}
