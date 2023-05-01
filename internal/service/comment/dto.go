package comment

type Create struct {
	DeveloperId int     `json:"developer_id"`
	TaskId      int     `json:"task_id"`
	Text        *string `json:"text"`
	CreatedBy   int     `json:"created_by"`
}

type Update struct {
	Id   int     `json:"id"`
	Text *string `json:"text"`
}

type List struct {
	Id          int     `json:"id"`
	DeveloperId int     `json:"developer_id"`
	TaskId      int     `json:"task_id"`
	Text        *string `json:"text"`
}

type Filter struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}
