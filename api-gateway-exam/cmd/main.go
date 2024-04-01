package main

import (
	"code-service-exam/api-gateway-exam/api"
	_ "code-service-exam/api-gateway-exam/api/docs" // swag
	"code-service-exam/api-gateway-exam/config"
	"code-service-exam/api-gateway-exam/pkg/logger"
	"code-service-exam/api-gateway-exam/services"

	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	// Postgres orqali Casbin role management ni nazorat qilish
	// Casbin implementation ACL with Postgres
	// psqlString := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
	// 	"localhost",
	// 	5432,
	// 	"admin",
	// 	"123",
	// 	"userdb")

	// db, err := gormAdapter.NewAdapter("postgres", psqlString, true)
	// if err != nil {
	// 	log.Error("gormadapter error", logger.Error(err))
	// }

	// enforcer, err := casbin.NewEnforcer("auth.conf", db)
	// if err != nil {
	// 	log.Error("NewForcer error", logger.Error(err))
	// }

	// Casbin implementation ACL with csv
	// a := fileadapter.NewAdapter("auth.csv")

	// bu yerda esa csv file orqali API larga dostup olishimiz mumkin bo'ladi
	enforcer, err := casbin.NewEnforcer("auth.conf", "auth.csv")
	if err != nil {
		log.Error("failed to create enforcer: %v", logger.Error(err))
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		log.Error("failed to load policy: %v", logger.Error(err))
	}

	enforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch", util.KeyMatch)
	enforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	// //in here we started the writer ...
	// writer, err := producer.NewKafkaProducerInit([]string{"localhost:9092"})
	// if err != nil {
	// 	log.Error("NewKafkaProducerInit: %v", logger.Error(err))
	// }

	// // in here we are producing message
	// err = writer.ProduceMessage("test-topic", []byte("Let's start messaging!"))
	// if err != nil {
	// 	log.Fatal("failed to run http server", logger.Error(err))
	// }

	// defer writer.Close()

	server := api.New(&api.Option{
		Conf:           cfg,
		Logger:         log,
		Enforcer:       enforcer,
		ServiceManager: serviceManager,
		// Writer:         writer,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

	r := gin.New()

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Serve Swagger JSON
	r.GET("/swagger.json", func(c *gin.Context) {
		c.File("docs/swagger.json")
	})

	// Your other routes...

}
