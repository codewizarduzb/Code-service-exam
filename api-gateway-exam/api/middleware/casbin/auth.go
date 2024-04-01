package casbin

import (
	"code-service-exam/api-gateway-exam/api/tokens"
	"code-service-exam/api-gateway-exam/config"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type CasbinHandler struct {
	config   config.Config
	enforcer *casbin.Enforcer
}

// Casbin Permission Checker Constructor
func CheckCasbinPermission(casbin *casbin.Enforcer, conf config.Config) gin.HandlerFunc {
	casnHandler := CasbinHandler{
		config:   conf,
		enforcer: casbin,
	}
	return func(ctx *gin.Context) {
		allowed, err := casnHandler.CheckPermission(ctx.Request)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
		} else if !allowed {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "check authorization",
			})
		}

	}
}

// the method that gets a role
func (casb *CasbinHandler) GetRole(ctx *http.Request) (string, error) {
	var t string

	token := ctx.Header.Get("Authorization")
	if token == "" {
		return "unauthorized", nil
	} else if strings.Contains(token, "Bearer") {
		t = strings.TrimPrefix(token, "Bearer ")
	} else {
		t = token
	}

	claims, err := tokens.ExtractClaim(t, []byte(config.Load().SignInKey))
	if err != nil {
		return "unauthorized, token is invalid", err
	}

	return cast.ToString(claims["role"]), nil
}

// permission checker
func (casb *CasbinHandler) CheckPermission(r *http.Request) (bool, error) {
	role, err := casb.GetRole(r)
	if err != nil {
		return false, err
	}

	method := r.Method
	action := r.URL.Path

	c, err := casb.enforcer.Enforce(role, action, method)
	if err != nil {
		return false, err
	}

	return c, nil

}
