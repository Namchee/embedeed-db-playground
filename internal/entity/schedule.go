package entity

import (
	"encoding/json"
)

type Schedule struct {
	ProductCode  string          `db:"product_code"`
	ScheduleCode string          `db:"schedule_code"`
	ScheduleName string          `db:"schedule_name"`
	Status       int             `db:"status"`
	Address      string          `db:"address"`
	Details      ScheduleDetails `db:"detail"`
}

type ScheduleDetail struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Time string `json:"time"`
}

type ScheduleDetails []ScheduleDetail

func (s *ScheduleDetails) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		_ = json.Unmarshal(v, &s)
		return nil
	case string:
		_ = json.Unmarshal([]byte(v), &s)
		return nil
	default:
		_ = json.Unmarshal([]byte("[]"), &s)
		return nil
	}
}
