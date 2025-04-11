package helper

import (
	"context"
	"encoding/json"

	jwt "github.com/dgrijalva/jwt-go"
)

const UserIDKey = "userid"

func GetJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[UserIDKey] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func GetJWTUserID(ctx context.Context) int64 {
	userid, _ := ctx.Value(UserIDKey).(json.Number).Int64()
	return userid
}
