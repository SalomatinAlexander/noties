package store

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: `toml:"database_url"`,
	}
}
