package entity

import "github.com/uptrace/bun"

type Project struct {
	bun.BaseModel `bun:"table:projects"`
	BasicEntity
	EmployeeId   int     `json:"employee_id" bun:"employee_id"`
	Name         *string `json:"name" bun:"name"`
	StartedDate  *string `json:"started_date" bun:"started_date"`
	FinishedDate *string `json:"finished_date" bun:"finished_date"`
	Status       *string `json:"status" bun:"status"`
	FileUrl      *string `json:"file_url" bun:"file_url"`
}
