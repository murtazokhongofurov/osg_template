package entity

import "time"

type BasicEntity struct {
	Id        int       `json:"id" bun:"id,pk,autoincrement"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
	CreatedBy int       `json:"created_by" bun:"created_by"`
	DeletedAt time.Time `json:"deleted_at" bun:"deleted_at"`
	DeletedBy int       `json:"deleted_by" bun:"deleted_by"`
}
