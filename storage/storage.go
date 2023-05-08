package storage

import (
	msql "github.com/dnevsky/hmm-bot/storage/mysql"

	"github.com/dnevsky/hmm-bot/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CheckUser(vkId int) (bool, error) // check user in db
}

type Find interface {
	GetFind() (models.Find, error)
}

type Storage struct {
	User
	Find
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		// postgr.NewUserPostgres(db),
		msql.NewUserMySQL(db),
		msql.NewFindMySQL(db),
	}
}
