package tdd_backend

type Config struct {
	BindAddr string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8082",
	}
}
