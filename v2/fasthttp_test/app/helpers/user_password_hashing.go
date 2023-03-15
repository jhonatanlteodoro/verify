package helpers

import "golang.org/x/crypto/bcrypt"

var hashCost = bcrypt.DefaultCost

func GenerateHashPassword(password string) (string, error) {
	bPass := []byte(password)
	hPass, err := bcrypt.GenerateFromPassword(bPass, hashCost)
	if err != nil {
		return "", err
	}

	hasedPassword := string(hPass)
	return hasedPassword, nil
}
