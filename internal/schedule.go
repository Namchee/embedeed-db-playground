package internal

import (
	"context"

	"github.com/Namchee/ramsql-playground/internal/entity"
)

type ScheduleRepository interface {
	GetSchedulesByProductID(
		ctx context.Context,
		productID int,
	) ([]entity.Schedule, error)
}
