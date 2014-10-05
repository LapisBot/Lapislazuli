package config

type Server struct {
	Login *Login `json:"login"`
	Nick *Nickname `json:"nick"`
	Messages *Messages `json:"messages"`
}

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
	NickServ *NickServ `json:"nickserv"`
}

type NickServ struct {
	User     string `json:"name"`
	Password string `json:"password"`
}

type Nickname struct {
	Name     string `json:"name"`
	Realname string `json:"realname"`
}

type Messages struct {
	Version string `json:"version"`
	Quit    string `json:"quit"`
}

func NewServer() *Server {
	return &Server {
		&Login {
			NickServ: &NickServ{},
		},
		&Nickname {},
		&Messages {},
	}
}
