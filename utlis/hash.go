package utlis

// this file is responcible for hashing data or password

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	ans, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(ans), err // in this string conversion we are converting byte slice to string
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}
