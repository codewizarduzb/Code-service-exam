package v1

import (
	"code-service-exam/api-gateway-exam/api/tokens"
	"code-service-exam/api-gateway-exam/config"

	"code-service-exam/api-gateway-exam/pkg/logger"
	"code-service-exam/api-gateway-exam/services"

	"github.com/casbin/casbin/v2"
)

type HandlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	jwtHandler     tokens.JWTHandler
	enforcer       *casbin.Enforcer
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	JWTHandler     tokens.JWTHandler
	Enforcer       *casbin.Enforcer
}

// New ...
func New(c *HandlerV1Config) *HandlerV1 {
	return &HandlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		jwtHandler:     c.JWTHandler,
		enforcer:       c.Enforcer,
	}
}
