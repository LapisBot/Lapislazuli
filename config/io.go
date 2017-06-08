package config

import (
	"encoding/json"
	"io"
)

func Parse(reader io.Reader) (*Config, error) {
	file, err := Read(reader)
	if err != nil {
		return nil, err
	}
	return file.Parse(), nil
}

func Read(reader io.Reader) (conf *ConfigFile, err error) {
	conf = New()
	err = json.NewDecoder(reader).Decode(conf)
	return
}

func Write(writer io.Writer, conf *Config) (err error) {
	// We can't use the encoder here becaus
func Write(writer io.Writer, conf *ConfigFile) (err error) {
	// We can't use the encoder here because it is not able to print indented
	result, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		return
	}
	_, err = writer.Write(result)
	if err != nil {
		return
	}
	_, err = io.WriteString(writer, "\n")
	return
}
