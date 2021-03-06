package client

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	ClientID string `koanf:"clientID"`
	Username string `koanf:"user"`
	Password string `koanf:"password"`
	Debug    bool   `koanf:"debug"`
	Topic    string `koanf:"topic"`
}
