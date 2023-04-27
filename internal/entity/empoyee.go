package entity

import "github.com/uptrace/bun"

type Employee struct {
	bun.BaseModel `bun:"table:employees"`
	BasicEntity
	FullName     *string `json:"full_name" bun:"full_name"`
	ProfilePhoto *string `json:"profile_photo" bun:"profile_photo"`
	Phone        *string `json:"phone" bun:"phone"`
	BirthDate    *string `json:"birth_date" bun:"birth_date"`
	Position     *string `json:"position" bun:"position"`
	Role         *string `json:"role" bun:"role"`
}
