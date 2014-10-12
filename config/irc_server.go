package config

type Server struct {
	Name string `json:"-"`
	Connection *Connection `json:"connection"`
	Login      *Login      `json:"login"`
	Messages   *Messages   `json:"messages"`
	Channels map[string]*Channel `json:"-"`
}

type Connection struct {
	Address     string `json:"address"`
	SSL         bool   `json:"ssl"`
	Certificate bool   `json:"certificate"`
}

type Login struct {
	User     string    `json:"user"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Ident    string    `json:"ident"`
	NickServ *NickServ `json:"nickserv"`
}

type NickServ struct {
	User     string `json:"name"`
	Password string `json:"password"`
}

type Messages struct {
	Version string `json:"version"`
	Quit    string `json:"quit"`
}

func NewServer(name string) *Server {
	return &Server{
		name,
		&Connection{
			SSL:         false,
			Certificate: true,
		},
		&Login{
			NickServ: &NickServ{},
		},
		&Messages{},
		nil,
	}
}
