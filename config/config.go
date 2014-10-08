package config

import "encoding/json"

type Config struct {
	Servers map[string]*Server `json:"servers"`
}

func New() *Config {
	return &Config{
		make(map[string]*Server, 0),
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
