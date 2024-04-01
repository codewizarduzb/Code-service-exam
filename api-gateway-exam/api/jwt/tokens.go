package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// JWTHandler ...
type JWTHandler struct {
	Sub     string
	Exp     string
	Iat     string
	Role    string
	Timeout int
}

type CustomClaims struct {
	*jwt.Token
	Sub  string  `json:"sub"`
	Exp  float64 `json:"exp"`
	Iat  float64 `json:"iat"`
	Role string  `json:"role"`
}

// GenerateAuthJWT ...
func (jwtHandler *JWTHandler) GenerateAuthJWT() (access string, err error) {
	var (
		accessToken *jwt.Token
		claims      jwt.MapClaims
	)

	accessToken = jwt.New(jwt.SigningMethodHS256)

	claims = accessToken.Claims.(jwt.MapClaims)
	claims["sub"] = jwtHandler.Sub
	claims["role"] = jwtHandler.Role
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(jwtHandler.Timeout)).Unix()
	claims["iat"] = time.Now().Unix()

	access, err = accessToken.SignedString([]byte("uzbekcodewizard"))
	if err != nil {
		fmt.Println("error while creating access token", err.Error())
		return
	}

	return
}

// ExtractClaim extracts claims from given token
func ExtractClaim(tokenStr string, signinigKey []byte) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return signinigKey, nil
	})
	if err != nil {
		return nil, err
	}
	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return claims, nil
}
