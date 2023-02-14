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
			product_code,
			schedule_code,
			schedule_name,
			start_date,
			status,
			address,
			detail
		FROM 
			product_kpk_schedule
		WHERE 
			status <> $1 AND
			product_id = $2
	`

	rows, err := s.db.QueryContext(ctx, query, constant.ScheduleInactive, productID)
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
			&s.StartDate,
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
