package entity

import "github.com/uptrace/bun"

type Attendance struct {
	bun.BaseModel `bun:"table:attendance"`
	BasicEntity
	EmployeeId *int    `json:"employee_id" bun:"employee_id"`
	Type       *string `json:"type" bun:"type"`
}
