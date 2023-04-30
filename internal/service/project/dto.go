package project

type Filter struct {
	Limit  *int
	Offset *int
	Search *string
}

type Create struct {
	EmployeeId   int     `json:"employee_id"`
	Name         *string `json:"name"`
	StartedDate  *string `json:"started_date"`
	FinishedDate *string `json:"finished_date"`
	FileUrl      *string `json:"file_url"`
	Status       *string `json:"status"`
	CreatedBy    int     `json:"created_by" `
}

type Update struct {
	Id           int     `json:"id"`
	Name         *string `json:"name"`
	StartedDate  *string `json:"started_date"`
	FinishedDate *string `json:"finished_date"`
	FileUrl      *string `json:"file_url"`
	Status       *string `json:"status"`
}

type List struct {
	Id           int     `json:"id"`
	Name         *string `json:"name"`
	StartedDate  *string `json:"started_date"`
	FinishedDate *string `json:"finished_date"`
	Status       *string `json:"status"`
}

type Detail struct {
	Id           int     `json:"id"`
	Name         *string `json:"name"`
	StartedDate  *string `json:"started_date"`
	FinishedDate *string `json:"finished_date"`
	FileUrl      *string `json:"file_url"`
	Status       *string `json:"status"`
	CreatedBy    int     `json:"created_by"`
}
