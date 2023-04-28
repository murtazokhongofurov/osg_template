package attendance

type Create struct {
	EmployeeId *int    `json:"employee_id"`
	Type       *string `json:"type"`
	CreatedBy  *string `json:"created_by"`
}

type Filter struct {
	Limit  *int    `json:"limit"`
	Offset *int    `json:"offset"`
	Search *string `json:"search"`
}

type Update struct {
	Id   int     `json:"id"`
	Type *string `json:"type"`
}

type List struct {
	Id         *int    `json:"id"`
	EmployeeId *int    `json:"employee_id"`
	Type       *string `json:"type"`
}

type Detail struct {
	Id         *int    `json:"id"`
	EmployeeId *int    `json:"employee_id"`
	Type       *string `json:"type"`
}
