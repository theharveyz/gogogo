package config

import (
	"encoding/json"
	"os"
	"sync"
)

type (
	ConfigIface interface {
		Register(file string) error
		Load(state interface{}) error
	}

	Config struct {
		lock *sync.RWMutex
		file string
	}
)

func New() *Config {
	return &Config{lock: new(sync.RWMutex)}
}

func (cfg *Config) Register(file string) error {
	_, error := os.Stat(file)

	if os.IsExist(error) {
		return error
	}
	// 注意 锁的粒度
	defer cfg.lock.Unlock()
	cfg.lock.Lock()
	cfg.file = file
	return nil
}

func (cfg *Config) Load(state interface{}) error {
	cfg.lock.RLock()
	fp, err := os.Open(cfg.file)
	cfg.lock.RUnlock() // 注意粒度

	if err != nil {
		return err
	}
	defer fp.Close()
	buf, bufs := make([]byte, 1024), []byte{} // 防止有0项
	for {
		n, _ := fp.Read(buf)
		if n == 0 {
			break
		}
		bufs = append(bufs, buf[:n]...)
	}
	err = json.Unmarshal(bufs, state)
	if err != nil {
		return err
	}
	return nil
}
