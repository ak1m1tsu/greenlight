package data

import (
	"database/sql"
	"errors"
)

const ContextCanceledByUser string = "pq: canceling statement due to user request"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
	Users  UserModel
	Tokens TokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db},
		Tokens: TokenModel{DB: db},
	}
}
