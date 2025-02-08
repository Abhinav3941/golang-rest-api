package utlis

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "mynameisanthonoy"

func GenerateJWT(Email string, User_id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": User_id,
		"email":   Email,
		"exp":     time.Now().Add(time.Hour * time.Duration(2)).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token,
		func(tokn *jwt.Token) (interface{}, error) {
			_, ok := tokn.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(secretKey), nil
		})

	if err != nil {
		return 0, errors.New("token parse error")
	}
	tokenisValid := parsedToken.Valid
	if !tokenisValid {
		return 0, errors.New("token is invalid")
	}
	//

	// this is the part of data extraction from the token

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("claims error")
	}
	//
	//email := claims["email"].(string)
	userid := int64(claims["user_id"].(float64))

	return userid, nil

}
