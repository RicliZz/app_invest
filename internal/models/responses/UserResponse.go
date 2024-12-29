package responses

import "github.com/RicliZz/app_invest/internal/models/userModel"

type UserResponse struct {
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	MiddleName          string `json:"middleName"`
	Email               string `json:"email"`
	Password            string `json:"password"`
	UserDetailsResponse `json:"userDetails"`
}

type UserDetailsResponse struct {
	Balance float64           `json:"balance"`
	Socials userModel.Socials `json:"socials"`
}
