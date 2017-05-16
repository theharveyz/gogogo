package config

import (
	"sync"
)

type (
	ConfigIface interface {
	}

	ConfigParser interface {
	}

	Config struct {
		File   string
		Parser ConfigParser
		mu     *sync.RWMutex
		items  map[interface{}]interface{}
	}
)

func New(parser ConfigParser) *Config {
	return &Config{mu: new(sync.RWMutex)}
}

func (conf *Config) Items() map[interface{}]interface{} {
	defer conf.mu.RUnlock()
	conf.mu.RLock()
	items := make(map[interface{}]interface{})
	for k, v := range conf.items {
		items[k] = v
	}
	return items
}
