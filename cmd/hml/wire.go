//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"example/hml/config"
	"example/hml/internal/server"
	"example/hml/internal/service"
	"github.com/Rascal0814/boot/kratos/depend"

	"github.com/Rascal0814/boot/kratos/log"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp() (*kratos.App, func(), error) {
	panic(wire.Build(
		log.ProvideLogger,
		depend.NewConsulRegistrar,
		server.ProviderSet,
		service.ProviderSet,
		config.ProviderSet,
		newApp))
}
