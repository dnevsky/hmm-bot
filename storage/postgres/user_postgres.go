package postgr

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CheckUser(vkId int) (bool, error) {
	var id int

	query := fmt.Sprintf("SELECT id FROM %s WHERE memberId = $1", hmmBotTable)
	row := r.db.QueryRow(query, vkId)
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		logrus.Printf("CheckUser failed: %s", err)
		return false, err
	}
	return true, nil
}
