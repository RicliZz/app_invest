package authModel

type RequestSignUpPayload struct {
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName"  validate:"required"`
	MiddleName string `json:"middleName" validate:"required"`
	Email      string `json:"email"     validate:"required,email"`
	Password   string `json:"password"  validate:"required,min=8"`
}

type RequestSignInPayload struct {
	Email    string `json:"email"    validate:"required,email"  example:"admin@mail.ru"`
	Password string `json:"password" validate:"required,min=8"  example:"12345678"`
}

type ResponseSignInPayload struct {
	Jwt string `json:"token"`
}
