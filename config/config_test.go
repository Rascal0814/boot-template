package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config := NewConfig()

	fmt.Print(config)
}
