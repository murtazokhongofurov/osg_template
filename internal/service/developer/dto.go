package developer

type Create struct {
	EmployeeId    int     `json:"employee_id"`
	DeveloperRole *string `json:"developer_role"`
	CreatedBy     int    `json:"created_by"`
}

type Update struct {
	Id            int     `json:"id"`
	DeveloperRole *string `json:"developer_role"`
}

type List struct {
	Id            int     `json:"id"`
	EmployeeId    int     `json:"employee_id"`
	DeveloperRole *string `json:"developer_role"`
}

type Detail struct {
	Id            int    `json:"id"`
	EmployeeId    int    `json:"employee_id"`
	DeveloperRole *string `json:"developer_role"`
}

type Filter struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}
