package vanilla

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Namchee/ramsql-playground/internal/entity"
	"github.com/stretchr/testify/assert"

	_ "github.com/proullon/ramsql/driver"
)

var (
	db *sql.DB
)

func TestMain(m *testing.M) {
	var err error

	db, err = sql.Open("ramsql", "TestScheduleVanillaDB")
	if err != nil {
		log.Fatalln("failed to open in-memory database: ", err)
	}

	queries := []string{
		`CREATE TABLE product_kpk_schedule
		(
			id bigserial,
			product_id bigint NOT NULL,
			product_code text NOT NULL,
			schedule_code text NOT NULL,
			schedule_name text NOT NULL,
			start_date date NOT NULL,
			end_date date NOT NULL,
			detail jsonb NOT NULL DEFAULT '[]',
			status smallint NOT NULL,
			address text NOT NULL,
			created_date timestamp with time zone NOT NULL,
			updated_date timestamp with time zone,
			expired_date timestamp with time zone NOT NULL,
			updated_by bigint NOT NULL
		);`,
		`INSERT INTO product_kpk_schedule
		(
			product_id,
			product_code,
			schedule_code,
			schedule_name,
			start_date,
			end_date,
			status,
			address,
			created_date,
			expired_date,
			updated_by
		)
		VALUES
		(
			1,
			'product-code',
			'schedule-code',
			'name',
			'2023-02-14',
			'2023-02-16',
			1,
			'a',
			'2016-06-22 19:10:25-07',
			'2024-06-23 19:10:25-07',
			1
		);`,
		`INSERT INTO product_kpk_schedule
		(
			product_id,
			product_code,
			schedule_code,
			schedule_name,
			start_date,
			end_date,
			status,
			address,
			created_date,
			expired_date,
			updated_by
		)
		VALUES
		(
			2,
			'product-code',
			'schedule-code',
			'name',
			'2023-02-14',
			'2023-02-16',
			1,
			'a',
			'2016-06-22 19:10:25-07',
			'2022-06-23 19:10:25-07',
			1
		)`,
		`INSERT INTO product_kpk_schedule
		(
			product_id,
			product_code,
			schedule_code,
			schedule_name,
			start_date,
			end_date,
			status,
			address,
			created_date,
			expired_date,
			updated_by
		)
		VALUES
		(
			5,
			'product-code',
			'schedule-code',
			'name',
			'2023-02-14',
			'2023-02-16',
			1,
			'a',
			'2016-06-22 19:10:25-07',
			'2024-06-23 19:10:25-07',
			1
		)`,
		`INSERT INTO product_kpk_schedule
		(
			product_id,
			product_code,
			schedule_code,
			schedule_name,
			start_date,
			end_date,
			status,
			address,
			created_date,
			expired_date,
			updated_by
		)
		VALUES
		(
			2,
			'product-code',
			'schedule-code',
			'name',
			'2023-02-14',
			'2023-02-16',
			2,
			'a',
			'2016-06-22 19:10:25-07',
			'2024-06-23 19:10:25-07',
			1
		)`,
	}

	for _, q := range queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Fatalln("failed to execute database query: ", err)
		}
	}

	code := m.Run()

	os.Exit(code)
}

func TestNewScheduleVanillaDB(t *testing.T) {
	assert.NotPanics(t, func() {
		NewScheduleVanillaDB(db)
	})
}

func TestScheduleVanillaRepository_GetSchedulesByProductID(
	t *testing.T,
) {
	want := []entity.Schedule{
		{
			ProductCode:  "product-code",
			ScheduleCode: "schedule-code",
			ScheduleName: "name",
			StartDate:    time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
			Address:      "a",
			Status:       1,
			Details:      entity.ScheduleDetails{},
		},
	}

	repo := &scheduleVanillaRepository{db: db}
	got, err := repo.GetSchedulesByProductID(context.Background(), 5)

	assert.Equal(t, want, got)
	assert.Equal(t, nil, err)
}
