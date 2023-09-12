package config

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/google/wire"
)

// ProviderSet 提供配置相关的依赖
var ProviderSet = wire.NewSet(
	NewConfig,
)

func NewConfig() *Config {
	c := config.New(
		config.WithSource(
			file.NewSource("config.yaml"),
		),
	)
	defer func() { _ = c.Close() }()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc Config
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	return &bc
}
