package msql

import (
	"VKBotAPI/models"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/jmoiron/sqlx"
)

type FindMySQL struct {
	db *sqlx.DB
}

func NewFindMySQL(db *sqlx.DB) *FindMySQL {
	return &FindMySQL{db: db}
}

func (r *FindMySQL) GetFind() (models.Find, error) {
	var find models.Find
	var pl string

	row := r.db.QueryRow("SELECT * FROM find WHERE id = 1") // там только одна такая запись, которая постоянно обновляется
	if err := row.Scan(&find.ID, &find.LS, &find.SF, &find.LV, &pl, &find.Client, &find.IP, &find.TS); err != nil {
		if err == sql.ErrNoRows {
			return find, errors.New("GetFind: find not found.")
		}
		return find, errors.New("GetFind: error while search find.")
	}

	if ok := json.Unmarshal([]byte(pl), &find.Players); ok != nil {
		return models.Find{}, errors.New("GetFind: error while parsing players string")
	}

	return find, nil

}
