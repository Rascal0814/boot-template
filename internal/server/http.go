package server

import (
	v1 "example/hml/api/v1"
	"example/hml/internal/service"
	"github.com/Rascal0814/boot/config"
	kratosmiddleware "github.com/Rascal0814/boot/kratos/middleware"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *config.Server, echo *service.EchoService, logger log.Logger) (*http.Server, error) {
	var opts = kratosmiddleware.DefaultHttpMiddleWare
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterEchoServiceHTTPServer(srv, echo)
	return srv, nil
}
