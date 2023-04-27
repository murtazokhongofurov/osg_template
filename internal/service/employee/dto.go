package employee

type Create struct {
	FullName     *string `json:"full_name" form:"full_name"`
	ProfilePhoto *string `json:"profile_photo" form:"profile_photo"`
	Phone        *string `json:"phone" form:"phone"`
	BirthDate    *string `json:"birth_date" form:"birth_date"`
	Position     *string `json:"position" form:"position"`
	Role         *string `json:"role" form:"role"`
	CreatedBy    int     `json:"-" form:"-"`
	Token        string  `json:"-" form:"-"`
}

type Update struct {
	Id           *int    `json:"id" form:"id"`
	FullName     *string `json:"full_name" form:"full_name"`
	ProfilePhoto *string `json:"profile_photo" form:"profile_photo"`
	Phone        *string `json:"phone" form:"phone"`
	BirthDate    *string `json:"birth_date" form:"birth_date"`
	Position     *string `json:"position" form:"position"`
	Role         *string `json:"role" form:"role"`
}

type List struct {
	Id           *int    `json:"id" form:"id"`
	FullName     *string `json:"full_name" form:"full_name"`
	ProfilePhoto *string `json:"profile_photo" form:"profile_photo"`
	Position     *string `json:"position" form:"position"`
	Role         *string `json:"role" form:"role"`
}

type Detail struct {
	Id           *int    `json:"id" form:"id"`
	FullName     *string `json:"full_name" form:"full_name"`
	ProfilePhoto *string `json:"profile_photo" form:"profile_photo"`
	Phone        *string `json:"phone" form:"phone"`
	BirthDate    *string `json:"birth_date" form:"birth_date"`
	Position     *string `json:"position" form:"position"`
	Role         *string `json:"role" form:"role"`
}

type Filter struct {
	Limit  *int    `json:"limit" form:"limit"`
	Offset *int    `json:"offset" form:"ofset"`
	Search *string `json:"search" form:"search"`
}
