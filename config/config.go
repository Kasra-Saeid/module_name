package config

import "os"

type Config struct {
	PG
	Credentials
}

type PG struct {
	URL string
}

type Credentials struct {
	SecretKey string
}

func New() *Config {
	pgUrl := os.Getenv("pg_url")
	secretKey := os.Getenv("secret_key")

	config := &Config{
		PG: PG{
			URL: pgUrl,
		},
		Credentials: Credentials{
			SecretKey: secretKey,
		},
	}
	return config
}
