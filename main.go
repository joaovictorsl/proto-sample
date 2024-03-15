package main

import (
	"net"

	"github.com/joaovictorsl/proto-sample/cmd/svc"
	"github.com/joaovictorsl/proto-sample/proto/user"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	listener net.Listener
	server   *grpc.Server
	logger   *zap.Logger
)

func main() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	initListener()

	server = grpc.NewServer()
	user.RegisterUserServiceServer(server, svc.NewUserServiceServer())

	logger.Info("Starting gRPC server ...")
	go func() {
		if err := server.Serve(listener); err != nil {
			logger.Panic("Failed to start gRPC server", zap.Error(err))
		}
	}()

	runClient()
}

func initListener() {
	var err error
	addr := "localhost:50051"

	listener, err = net.Listen("tcp", addr)
	if err != nil {
		logger.Panic("Failed to listen", zap.Error(err))
	}

	logger.Info("Listening on " + addr)
}
