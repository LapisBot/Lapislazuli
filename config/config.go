package config

import "encoding/json"

type Config struct {
	Servers  map[string]*Server  `json:"servers"`
	Channels map[string]*Channel `json:"channels"`
}

func New() *Config {
	return &Config{
		make(map[string]*Server),
		make(map[string]*Channel),
	}
}

func (conf *Config) String() string {
	buf, err := json.Marshal(conf)
	if err == nil {
		return string(buf)
	} else {
		return err.Error()
	}
}
