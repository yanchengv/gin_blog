package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SigningKey []byte
}

//token签名内携带的信息
type MyCustomClaims struct {
	ID       uint
	Nickname string
	jwt.StandardClaims
}

type JWTCustomInfo struct {
	ID       uint
	Nickname string
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("秘钥key"),
	}
}

//生成jwt的token
// j := &utils.JWT{}
// j.CreateToken(user)
func (j *JWT) CreateToken(user JWTCustomInfo) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour) //60 * time.Second，过期时间
	issuer := "yan"                          //签发者
	claims := MyCustomClaims{
		ID:       user.ID,
		Nickname: user.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
	return token, err
}

//解析token
// j := &utils.JWT{}
// j.ParseToken(token)
func (j *JWT) ParseToken(token string) (*MyCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
