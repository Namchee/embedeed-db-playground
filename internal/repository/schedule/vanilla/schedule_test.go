package vanilla

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Namchee/ramsql-playground/internal/entity"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func TestMain(m *testing.M) {
	postgres := embeddedpostgres.NewDatabase()
	err := postgres.Start()
	if err != nil {
		log.Fatalln("failed to start in-memory database: ", err)
	}

	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln("failed to connect to in-memory database: ", err)
	}

	seedDatabase()

	code := m.Run()

	err = postgres.Stop()
	if err != nil {
		log.Fatalln("failed to stop in-memory database: ", err)
	}

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

func seedDatabase() {
	queries := []string{
		`CREATE TABLE product_kpk_schedule
		(
			id bigserial,
			product_id bigint NOT NULL,
			product_code character varying NOT NULL,
			schedule_code character varying NOT NULL,
			schedule_name character varying NOT NULL,
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
		_, err := db.Exec(q)
		if err != nil {
			log.Fatalln("failed to execute database query: ", err)
		}
	}
}
