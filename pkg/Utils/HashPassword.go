package Utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hash string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), password)
	return err == nil
}
