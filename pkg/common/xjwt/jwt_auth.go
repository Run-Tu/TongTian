package xjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/google/uuid"
)

func CreateToken(uid int64, platfrom int8, access bool, ex int) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "lark.com"
	now := time.Now()
	claims["exp"] = now.Add(time.Duration(ex) * time.Second)
	claims["iat"] = now.Unix()
	claims["session_id"] = uuid.New().String()
	claims["udi"] = uid 
}
