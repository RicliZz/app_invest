package authRepository

import (
	"database/sql"
)

type AuthRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db}
}

func (rep *AuthRepositoryImpl) Create() {

}
