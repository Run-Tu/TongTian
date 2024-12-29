// HMAC系列(对称加密) SigningMethodHS256, SigningMethodHS384, SigningMethodHS512，需要传入 []byte 类型的 共享密钥 来签名和验签。
package xjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/google/uuid"
)

var (
	jwt_key = []byte("runtu")
)

func CreateToken(uid int64, platfrom int8, access bool, ex int) (tk string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "lark.com"
	now := time.Now()
	claims["exp"] = now.Add(time.Duration(ex) * time.Second)
	claims["iat"] = now.Unix()
	claims["session_id"] = uuid.New().String()
	claims["uid"] = uid
	claims["platfrom"] = platfrom

	tokenStr, err := token.SignedString(jwt_key)

	return tokenStr, err
}
