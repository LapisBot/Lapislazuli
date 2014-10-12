package config

type Channel struct {
	Name string `json:"-"`
}

func NewChannel() *Channel {
	return &Channel{}
}
