package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pswd string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pswd), 14)
	if err != nil {
		return "", err
	}

	return string(hashed), err
}

func CheckPassword(pswd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
	return err == nil
}
