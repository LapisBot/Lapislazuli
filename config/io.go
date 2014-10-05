package config

import (
	"io"
	"encoding/json"
)

func Read(reader io.Reader) (conf *Config, err error) {
	conf = New()
	err = json.NewDecoder(reader).Decode(conf)
	return
}

func Write(writer io.Writer, conf *Config) (err error) {
	// We can't use the encoder here because it is not able to print indented
	result, err := json.MarshalIndent(conf, "", "  "); if err != nil { return }
	_, err = writer.Write(result); if err != nil { return }
	_, err = io.WriteString(writer, "\n")
	return
}
