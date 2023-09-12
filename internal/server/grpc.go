package server

import (
	v1 "example/hml/api/v1"
	"example/hml/internal/service"

	"github.com/Rascal0814/boot/config"
	kratosmiddleware "github.com/Rascal0814/boot/kratos/middleware"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *config.Server, echo *service.EchoService, logger log.Logger) (*grpc.Server, error) {
	var opts = kratosmiddleware.DefaultGrpcMiddleWare
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterEchoServiceServer(srv, echo)
	return srv, nil
}
