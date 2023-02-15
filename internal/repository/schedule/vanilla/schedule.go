package vanilla

import (
	"context"
	"database/sql"

	"github.com/Namchee/ramsql-playground/internal"
	"github.com/Namchee/ramsql-playground/internal/constant"
	"github.com/Namchee/ramsql-playground/internal/entity"
)

type scheduleVanillaRepository struct {
	db *sql.DB
}

func NewScheduleVanillaDB(
	db *sql.DB,
) internal.ScheduleRepository {
	return &scheduleVanillaRepository{
		db: db,
	}
}

func (s *scheduleVanillaRepository) GetSchedulesByProductID(
	ctx context.Context,
	productID int,
) ([]entity.Schedule, error) {
	result := []entity.Schedule{}

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

	rows, err := s.db.QueryContext(
		ctx,
		query,
		constant.ScheduleInactive,
		productID,
	)
	if err != nil {
		return result, nil
	}
	defer rows.Close()

	for rows.Next() {
		var s entity.Schedule

		err := rows.Scan(
			&s.ProductCode,
			&s.ScheduleCode,
			&s.ScheduleName,
			&s.Status,
			&s.Address,
			&s.Details,
		)
		if err != nil {
			return []entity.Schedule{}, nil
		}

		result = append(result, s)
	}

	return result, nil
}
