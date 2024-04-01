package api

import (
	_ "code-service-exam/api-gateway-exam/api/docs" // swag
	v1 "code-service-exam/api-gateway-exam/api/handlers/v1"
	casbinC "code-service-exam/api-gateway-exam/api/middleware/casbin"
	"code-service-exam/api-gateway-exam/config"
	"code-service-exam/api-gateway-exam/pkg/logger"
	"code-service-exam/api-gateway-exam/services"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	Enforcer       *casbin.Enforcer
	ServiceManager services.IServiceManager
}

// @Title Special for fourth exam
// @Version 1.0
// @Description Microservices
// @Host localhost:8080
func New(option *Option) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Enforcer:       option.Enforcer,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	api := router.Group("/v1")

	// bu yerdan Casbin orqali API larga dostup berishimiz mumkin
	api.Use(casbinC.CheckCasbinPermission(option.Enforcer, option.Conf))

	// User Service APIs
	api.POST("/createuser", handlerV1.CreateUser)
	api.GET("/getuser", handlerV1.GetUser)
	api.GET("/listusers", handlerV1.ListUsers)
	api.PATCH("/updateuser", handlerV1.UpdateUser)
	api.DELETE("/deleteuser", handlerV1.DeleteUser)

	// Post Service APIs
	api.POST("/createpost", handlerV1.CreatePost)
	api.GET("/getpostbyid", handlerV1.GetPostById)
	api.GET("/getpostsbyuserid", handlerV1.GetPostsByUserId)
	api.PATCH("/updatepostcontent", handlerV1.UpdatePost)
	api.DELETE("/deletepost", handlerV1.DeletePost)

	// Comment Service APIs
	api.POST("/createcomment", handlerV1.CreateComment)
	api.GET("/getcomment", handlerV1.GetComment)
	api.PATCH("/updatecomment", handlerV1.UpdateComment)
	api.DELETE("/deletecomment", handlerV1.DeleteComment)
	api.GET("/listcomments", handlerV1.ListComments)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
