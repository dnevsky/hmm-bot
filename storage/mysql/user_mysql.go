package msql

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserMySQL struct {
	db *sqlx.DB
}

func NewUserMySQL(db *sqlx.DB) *UserMySQL {
	return &UserMySQL{db: db}
}

func (r *UserMySQL) CheckUser(vkId int) (bool, error) {
	var id int

	query := fmt.Sprintf("SELECT id FROM %s WHERE memberId = ?", hmmBotTable)
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
