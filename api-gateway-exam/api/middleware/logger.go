package middleware

import (
	"code-service-exam/api-gateway-exam/api/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Authentification
func Auth(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/jwt" {
		ctx.Next()
		return
	}

	token := ctx.GetHeader("Authorization")

	// token bo'sh bo'lsa ham keyingi etapga o'tib ketadi
	if token == "" {
		ctx.Next()
		return
	}

	//token xato kiritilsayam keyingi etapga o'tadi
	claims, err := jwt.ExtractClaim(token, []byte("uzbekcodewizard"))
	if err != nil {
		ctx.Next()
		return
	}

	if cast.ToString(claims["role"]) != "user" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "no access to this path",
		})
		return
	}
	ctx.Next()
}
