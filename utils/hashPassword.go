package utils

import "golang.org/x/crypto/bcrypt"

func HashGenerateFromPassword(password *string) ([]byte, error) {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), cost)
	if err != nil {
		return nil, err
	}
	return hash, err
}
