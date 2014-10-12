package config

import (
	"encoding/json"
	"strings"
)

type ConfigFile struct {
	Servers  map[string]*Server  `json:"servers"`
	Channels map[string]*Channel `json:"channels"`
}

type Config struct {
	Servers map[string]*Server
}

func New() *ConfigFile {
	return &ConfigFile{
		make(map[string]*Server),
		make(map[string]*Channel),
		}
	}

func (conf *ConfigFile) Parse() *Config {
	parsed := &Config{make(map[string]*Server)}
	var def *Server

	for name, serverConf := range conf.Servers {
		server := &Server{}
		*server = *serverConf
		server.Name = name
		server.Channels = make(map[string]*Channel)
		parsed.Servers[name] = server

		if def == nil {
			def = server
		}
	}

	for name, _ := range conf.Channels {
		for _, name = range strings.Split(name, ",") {
			descriptor := strings.SplitN(name, ":", 2)
			name = descriptor[0]
			server := def
			if len(descriptor) == 2 {
				server = parsed.Servers[descriptor[0]]
				name = descriptor[1]
			}

			channel := server.Channels[name]
			if channel == nil {
				channel = &Channel{name}
				server.Channels[name] = channel
			}

			// TODO: Add configuration
		}
	}

	return parsed
}

func (conf *ConfigFile) String() string {
	buf, err := json.Marshal(conf)
	if err == nil {
		return string(buf)
	} else {
		return err.Error()
	}
}
