package task

type Filter struct {
	Limit  *int
	Offset *int
	Search *string
}

type Create struct {
	DeveloperId *int    `json:"developer_id" form:"developer_id"`
	Title       *string `json:"title" form:"title"`
	Description *string `json:"description" form:"description"`
	FileUrl     *string `json:"file_url" form:"file_url"`
	CreatedBy   *int    `json:"-" form:"-"`
	Token       *string `json:"-" form:"-"`
}

type Update struct {
	Id          int     `json:"id" form:"id"`
	Title       *string `json:"title" form:"title"`
	Description *string `json:"description" form:"description"`
	FileUrl     *string `json:"file_url" form:"file_url"`
}

type List struct {
	Id          int     `json:"id"`
	DeveloperId *int    `json:"developer_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	FileUrl     *string `json:"file_url"`
}

type Detail struct {
	Id          int     `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	FileUrl     *string `json:"file_url"`
}
