package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24 * 30

var mySecret = []byte("image")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (accessToken, refreshToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID, // 自定义字段
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "image",                                    // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	/*
		jwt.SigningMethodES256 两种类型 *SigningMethodECDSA 和 *SigningMethodHMAC
		jwts := jwt.NewWithClaims(jwt.SigningMethodES256, c) // SigningMethodES256 *SigningMethodECDSA  此类型会报错： key is of invalid type`
		jwts := jwt.NewWithClaims(jwt.SigningMethodHS256, c) // SigningMethodHS256 *SigningMethodHMAC 不报错
	*/
	// 草 SigningMethodHS256不报错 SigningMethodES256 报错
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "image",
	}).SignedString(mySecret)

	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(accessToken, refreshToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(refreshToken, keyFunc); err != nil {
		return
	}

	// 从旧access token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(accessToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID, claims.Username)
	}
	return
}
