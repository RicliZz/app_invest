package requests

type UpdateUserRequest struct {
	FirstName  *string `gorm:"column:firstName"         json:"firstName"`
	LastName   *string `gorm:"column:lastName"          json:"lastName"`
	MiddleName *string `gorm:"column:middleName"        json:"middleName"`
	Email      *string `json:"email"`
}

type UpdateUserDetailsRequest struct {
	Balance *float64 `json:"balance"`
}

type PatchProfileRequest struct {
	UpdateUserRequest    UpdateUserRequest        `json:"user"`
	UpdateDetailsRequest UpdateUserDetailsRequest `json:"user_details"`
}
