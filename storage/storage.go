package storage

import "github.com/jmoiron/sqlx"

type User interface {
	CheckUser(vkId int) (bool, error) // check user in db
}

type Storage struct {
	User
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		NewUserPostgres(db),
	}
}
