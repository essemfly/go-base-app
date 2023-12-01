package postgres

import (
	"essemfly/go_base_app/internal/domain"

	"gorm.io/gorm"
)

type postgresLogRepo struct {
	db *gorm.DB
}

func PostgresLogRepository(db *gorm.DB) LogRepository {
	return &postgresLogRepo{db: db}
}

func (r *postgresLogRepo) ListLogs() []*domain.LogData {
	var logs []*domain.LogData
	r.db.Find(&logs)
	return logs
}

func (r *postgresLogRepo) WriteLog(log *domain.LogData) error {
	return r.db.Create(log).Error
}
