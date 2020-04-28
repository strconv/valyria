package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/strconv/valyria/config"
)

// 参考 https://www.liwenzhou.com/posts/Go/jwt_in_gin/
const HEADER_TOKEN_KEY = "jwt_token" // http请求体内的token

var (
	TokenExpired     = errors.New("token expired")
	TokenNotValidYet = errors.New("token not valid yet")
	TokenMalformed   = errors.New("token malformed")
	TokenInvalid     = errors.New("invalid token")
	Secret           = []byte(config.C.JWT.Secret)
)

type Claims struct {
	UID int64 // 这里根据需要修改类型，考虑改为interface{}
	jwt.StandardClaims
}

func GenToken(uid int64) (string, error) {
	c := &Claims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.C.JWT.Timeout * time.Minute).Unix(),
			Issuer:    config.C.Service.Name,
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return Secret, nil
		},
	)
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			}
			if v.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			}
			if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			}
			return nil, TokenInvalid
		}
	}

	if token == nil {
		return nil, TokenInvalid
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
