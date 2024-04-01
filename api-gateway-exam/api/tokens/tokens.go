package tokens

import (
	"code-service-exam/api-gateway-exam/pkg/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTHandler struct {
	Sub       string
	Exp       string
	Iat       string
	Role      string
	SignInKey string
	Log       logger.Logger
	Token     string
	Timeout   int
}

type CustomClaims struct {
	*jwt.Token
	Sub  string  `json:"sub"`
	Exp  float64 `json:"exp"`
	Iat  float64 `json:"iat"`
	Role string  `json:"role"`
}

// Generating AuthGWT ...
func (jwtHander *JWTHandler) GenerateAuthJWT() (access, refresh string, err error) {
	var (
		accessToken  jwt.Token
		refreshToken jwt.Token
		claims       jwt.MapClaims
		rtClaims     jwt.MapClaims
	)

	accessToken = *jwt.New(jwt.SigningMethodHS256)
	refreshToken = *jwt.New(jwt.SigningMethodHS256)
	claims = accessToken.Claims.(jwt.MapClaims)
	claims["sub"] = jwtHander.Sub
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(jwtHander.Timeout)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = jwtHander.Role

	access, err = accessToken.SignedString([]byte(jwtHander.SignInKey))
	if err != nil {
		jwtHander.Log.Fatal("error generating refresh token")
		return
	}

	rtClaims = refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = jwtHander.Sub
	refresh, err = refreshToken.SignedString([]byte(jwtHander.SignInKey))
	if err != nil {
		jwtHander.Log.Fatal("error generating refresh token")
		return
	}
	return
}

// Extracting claims
func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	token, err = jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtHandler.SignInKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		jwtHandler.Log.Fatal("invalid jwt token")
		return nil, err
	}
	return claims, nil
}

// Extracting a claim
func ExtractClaim(tokenStr string, signinKey []byte) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	token, err = jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return signinKey, nil
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
