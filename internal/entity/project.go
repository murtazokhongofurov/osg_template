package entity

import "github.com/uptrace/bun"

type Prject struct {
	bun.BaseModel `bun:"table:prjects"`
	BasicEntity
	EmployeeId   *int    `json:"employee_id" bun:"employee_id"`
	Name         *string `json:"name" bun:"name"`
	StartedDate  *string `json:"started_date" bun:"started_date"`
	FinishedDate *string `json:"finished_date" bun:"finished_date"`
	Status       *string `json:"status" bun:"status"`
	FileUrl      *string `json:"file_url" bun:"file_url"`
}
