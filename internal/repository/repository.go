package repository

import "github.com/RicliZz/app_invest/internal/models/authModel"

type AuthRepository interface {
	Create(payload authModel.RequestSignUpPayload) error
}

type UserRepository interface {
	FindUserByEmail(email string) bool
}
