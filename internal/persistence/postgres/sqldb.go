package postgres

import (
	"essemfly/go_base_app/config"
	"essemfly/go_base_app/internal/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLDatabase struct {
	Repos *SQLRepository
}

func NewSQLDatabase(cfg config.Config) (*SQLDatabase, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbURL, cfg.DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.LogData{})
	db.AutoMigrate(&domain.LogAnalysisResult{})

	sqlRepository := InitPostgres(db)

	return &SQLDatabase{
		Repos: sqlRepository,
	}, nil
}

func InitPostgres(db *gorm.DB) *SQLRepository {
	return &SQLRepository{
		LogRepository:          PostgresLogRepository(db),
		LogAnalyticsRepository: nil,
	}
}
