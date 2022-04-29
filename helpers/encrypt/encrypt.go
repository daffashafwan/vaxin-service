package encrypt

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}