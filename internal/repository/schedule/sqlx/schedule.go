package sqlx

import (
	"context"

	"github.com/Namchee/ramsql-playground/internal"
	"github.com/Namchee/ramsql-playground/internal/constant"
	"github.com/Namchee/ramsql-playground/internal/entity"
	"github.com/jmoiron/sqlx"
)

type scheduleSQLXRepository struct {
	db *sqlx.DB
}

func NewScheduleSQLXDB(
	db *sqlx.DB,
) internal.ScheduleRepository {
	return &scheduleSQLXRepository{
		db: db,
	}
}

func (s *scheduleSQLXRepository) GetSchedulesByProductID(
	ctx context.Context,
	productID int,
) ([]entity.Schedule, error) {
	var result []entity.Schedule

	query := `
		SELECT 
			pks.product_code,
			pks.schedule_code,
			pks.schedule_name,
			pks.status,
			pks.address,
			pks.detail
		FROM 
			product_kpk_schedule pks
		WHERE 
			pks.status <> $1 AND
			pks.product_id = $2 AND
			pks.expired_date > NOW()
	`

	err := s.db.SelectContext(ctx, &result, query, constant.ScheduleInactive, productID)

	return result, err
}
