package storage

import (
	"VKBotAPI/models"
	msql "VKBotAPI/storage/mysql"

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
