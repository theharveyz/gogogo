package config

import (
	"fmt"
	"testing"
)

var cfg *Config = New()

func TestRegister(t *testing.T) {
	err := cfg.Register("./config.json")
	fmt.Println(err)
}

func TestLoad(t *testing.T) {
	cfg.Register("./config.json")
	state := &ConfigStatement{}
	err := cfg.Load(state)
	fmt.Println(state.Database.Default.Db, err)
}
