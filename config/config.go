package config

import (
	bc "github.com/Rascal0814/boot/config"
	"github.com/google/wire"
)

// ProviderSet 提供配置相关的依赖
var ProviderSet = wire.NewSet(
	NewConfig,
	wire.FieldsOf(new(*bc.Config), "Server", "Data"),
)

func NewConfig() (*bc.Config, error) {
	return bc.LoadConfig()
}
