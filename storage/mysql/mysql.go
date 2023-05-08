package msql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	hmmBotTable = "hmmBot"
)

type Config struct {
	Username string
	Password string
	Net      string
	Host     string
	DBName   string
}

func NewMySQLDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s?multiStatements=true",
		cfg.Username, cfg.Password, cfg.Net, cfg.Host, cfg.DBName))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
