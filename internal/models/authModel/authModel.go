package authModel

type RequestSignUpPayload struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type RequestSignInPayload struct {
	email    string
	password string
}

type ResponseSignInPayload struct {
	jwt string
}
