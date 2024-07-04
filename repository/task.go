package repository

import "github.com/jmoiron/sqlx"

type TaskDB struct {
	db *sqlx.DB
}

func NewTaskDB(db *sqlx.DB) *TaskDB {
	return &TaskDB{db: db}
}
