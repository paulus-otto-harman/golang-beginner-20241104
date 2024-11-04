package repository

import (
	"20241104/class/2/model"
	"database/sql"
)

type User struct {
	Db *sql.DB
}

func InitUserRepo(db *sql.DB) *User {
	return &User{Db: db}
}

func (repo *User) Create(user *model.User) (*model.Session, error) {
	query := `INSERT INTO users (fullname, email, password) VALUES ($1, $2, $3) RETURNING id`

	tx, err := repo.Db.Begin()
	if err != nil {
		return nil, err
	}

	if err = repo.Db.QueryRow(query, user.Fullname, user.Email, user.Password).Scan(&user.Id); err != nil {
		return nil, err
	}

	var session model.Session

	query = `INSERT INTO sessions (user_id) VALUES ($1) RETURNING id`
	if err = repo.Db.QueryRow(query, user.Id).Scan(&session.Id); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &session, nil
}
