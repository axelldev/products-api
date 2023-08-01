package server

type ServerConfig struct {
	Addr string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Addr: ":8080",
	}
}
