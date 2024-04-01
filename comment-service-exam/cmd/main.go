package main

import (
	"code-service-exam/comment-service-exam/config"
	pbc "code-service-exam/comment-service-exam/genproto/comment-proto"
	"code-service-exam/comment-service-exam/pkg/db"
	"code-service-exam/comment-service-exam/pkg/logger"
	"code-service-exam/comment-service-exam/service"
	grpcclient "code-service-exam/comment-service-exam/service/grpc_client"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "code-service-exam/comment-service-exam")

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	// in here we are connecting to Postgres
	connDB, err, _ := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	// we are dialing to another microservices
	grpcClient, err := grpcclient.New(cfg)
	if err != nil {
		log.Fatal("grpc client dail error", logger.Error(err))
	}

	// just implementation of user-service
	CommentService := service.NewCommentService(connDB, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pbc.RegisterCommentServiceServer(s, CommentService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
