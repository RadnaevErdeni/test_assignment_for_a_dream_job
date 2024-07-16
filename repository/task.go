package repository

import (
	"fmt"
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
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description,date_create,status) values ($1, $2,Now(),'not_started') RETURNING id", taskTable)
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
	query := fmt.Sprintf("SELECT ts.id,ts.title,ts.description,ts.start_time,ts.end_time,ts.duration,ts.total_pause_duration,ts.status,ts.last_resume_time,ts.last_pause_time, ts.date_create FROM %s ts INNER JOIN %s ut ON ts.id = ut.task_id WHERE ut.user_id = $1", taskTable, usersTaskTable)
	if err := r.db.Select(&tasks, query, userId); err != nil {
		return nil, err
	}
	return tasks, nil
}
func (r *TaskDB) GetById(userId, taskId int) (testtask.Tasks, error) {
	var user testtask.Tasks
	query := fmt.Sprintf(`SELECT ts.id,ts.title,ts.description,ts.start_time,ts.end_time,ts.duration,ts.total_pause_duration,ts.status,ts.last_resume_time,ts.last_pause_time, ts.date_create FROM %s ts INNER JOIN user_task us ON us.task_id = ts.id WHERE ts.id = $1 AND us.user_id = $2`,
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

func (r *TaskDB) Start(userId, taskId int) error {
	startQuery := fmt.Sprintf("UPDATE %s ts SET start_time = NOW(), last_resume_time = NOW(),status = 'in_progress' FROM %s ut WHERE ts.id = $1 AND ut.task_id = ts.id AND ut.user_id = $2 AND status = 'not_started'", taskTable, usersTaskTable)
	_, err := r.db.Exec(startQuery, taskId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskDB) End(userId, taskId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	endQuery := fmt.Sprintf(`UPDATE %s ts SET end_time = NOW(), status = 'completed' FROM %s ut WHERE ts.id = $1 AND ut.task_id = ts.id AND ut.user_id = $2 AND (status = 'in_progress' OR status = 'paused')`, taskTable, usersTaskTable)
	_, err = tx.Exec(endQuery, taskId, userId)
	if err != nil {
		return err
	}

	updateDurationQuery := fmt.Sprintf(`UPDATE %s ts SET duration = (end_time - start_time) - total_pause_duration FROM %s ut WHERE ts.id = $1 AND ut.task_id = ts.id AND ut.user_id = $2 AND status = 'completed'`, taskTable, usersTaskTable)
	_, err = tx.Exec(updateDurationQuery, taskId, userId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *TaskDB) Pause(userId, taskId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	pauseQuery := fmt.Sprintf(`UPDATE %s ts SET last_pause_time = NOW(), status = 'paused' FROM %s ut WHERE ts.id = $1 AND ut.task_id = ts.id AND ut.user_id = $2 AND status = 'in_progress'`, taskTable, usersTaskTable)
	_, err = tx.Exec(pauseQuery, taskId, userId)
	if err != nil {
		return err
	}
	updateDurationQuery := fmt.Sprintf(`UPDATE %s ts SET total_pause_duration = total_pause_duration + (NOW() - last_resume_time) FROM %s ut WHERE ts.id = $1 AND ut.task_id = ts.id AND ut.user_id = $2 AND status = 'paused'`, taskTable, usersTaskTable)
	_, err = tx.Exec(updateDurationQuery, taskId, userId)
	if err != nil {
		return err
	}

	return tx.Commit()

}

func (r *TaskDB) Resume(userId, taskId int) error {
	query := fmt.Sprintf(`UPDATE %s ts SET last_resume_time = NOW(),status = 'in_progress' FROM %s ut WHERE ts.id = $1 AND ut.task_id = ts.id AND ut.user_id = $2 AND status = 'paused'`, taskTable, usersTaskTable)
	_, err := r.db.Exec(query, taskId, userId)
	return err
}
