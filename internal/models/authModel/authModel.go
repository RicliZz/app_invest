package authModel

type RequestSignUpPayload struct {
	email    string
	password string
}

type RequestSignInPayload struct {
	email    string
	password string
}

type ResponseSignInPayload struct {
	jwt string
}
