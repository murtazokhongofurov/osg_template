package entity

import "github.com/uptrace/bun"

type Task struct {
	bun.BaseModel `bun:"table:tasks"`
	BasicEntity

	DeveloperId int     `json:"developer_id" bun:"developer_id"`
	Title       *string `json:"title" bun:"title"`
	Description *string `json:"description" bun:"description"`
	FileUrl     *string `json:"file_url" bun:"file_url"`
}
