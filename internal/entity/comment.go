package entity

import "github.com/uptrace/bun"

type Comment struct {
	bun.BaseModel `bun:"table:comment"`
	BasicEntity
	DeveloperId int     `json:"developer_id" bun:"developer_id"`
	TaskId      int     `json:"task_id" bun:"task_id"`
	Text        *string `json:"text" bun:"text"`
}
