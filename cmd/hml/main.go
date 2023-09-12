package main

import (
	"github.com/go-kratos/kratos/v2/registry"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name  string
	id, _ = os.Hostname()
)

func newApp(logger log.Logger, r registry.Registrar, gs *grpc.Server, hs *http.Server) (*kratos.App, error) {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(r),
	), nil
}

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}
	defer func() { cleanup() }()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
