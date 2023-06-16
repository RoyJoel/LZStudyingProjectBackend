package middleware

import (
	"errors"
	"fmt"

	// "net/http"
	"time"

	// "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	SecretKey = []byte("TennisMoment")
	Issuer    = "TennisMoment"
	ExpireAt  = time.Hour * 24 * 60
)

type JsonWebTokenClaim struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateToken(loginName string, password string) (string, error) {
	claims := JsonWebTokenClaim{
		loginName,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpireAt)),
			Issuer:    Issuer, // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func ParseToken(token string) (*JsonWebTokenClaim, error) {
	t, err := jwt.ParseWithClaims(token, new(JsonWebTokenClaim), func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := t.Claims.(*JsonWebTokenClaim); ok && t.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuth 鉴权中间件
// func JWTAuth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// 获取请求头中 token，实际是一个完整被签名过的 token；a complete, signed token
// 		tokenStr := c.GetHeader("Authorization")
// 		if tokenStr == "" {
// 			c.JSON(http.StatusForbidden, "No token. You don't have permission!")
// 			c.Abort()
// 			return
// 		}

// 		// 解析拿到完整有效的 token，里头包含解析后的 3 segment
// 		token, err := ParseToken(tokenStr)
// 		if err != nil {
// 			c.JSON(http.StatusForbidden, "Invalid token! You don't have permission!")
// 			c.Abort()
// 			return
// 		}

// 		// 将 claims 中的用户信息存储在 context 中
// 		c.Set("uid", token.Uid)
// 		c.Set("username", token.Username)

// 		// 这里执行路由 HandlerFunc
// 		c.Next()
// 	}
// }
