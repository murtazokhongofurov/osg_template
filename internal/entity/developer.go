package entity

import "github.com/uptrace/bun"

type Developer struct {
	bun.BaseModel `bun:"table:developers"`
	BasicEntity
	EmployeeId    *string `json:"employee_id" bun:"employee_id"`
	DeveloperRole *string `json:"developer_role" bun:"developer_role"`
}
