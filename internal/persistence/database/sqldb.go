package database

import (
	"essemfly/go_base_app/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLDatabase struct {
	db *gorm.DB
}

func (d *SQLDatabase) QuerySomething() string {
	return "something"
}

func NewSQLDatabase(cfg config.Config) (*SQLDatabase, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbURL, cfg.DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &SQLDatabase{db: db}, nil
}
