package config

import "encoding/json"

type Config struct {
	Servers []*Server `json:"servers"`
}

func New() *Config {
	return &Config{
		make([]*Server, 0),
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
