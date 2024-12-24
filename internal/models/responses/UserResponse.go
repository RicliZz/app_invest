package responses

type UserResponse struct {
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	MiddleName          string `json:"middleName"`
	Email               string `json:"email"`
	Password            string `json:"password"`
	UserDetailsResponse `json:"userDetails"`
}

type UserDetailsResponse struct {
	Balance float64 `json:"balance"`
}
