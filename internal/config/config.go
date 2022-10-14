package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Listen struct {
		BindIP string `env:"BIND_IP" env-default:"127.0.0.1"`
		Port   string `env:"SERVER_PORT" env-default:"8000"`
	}
	AppConfig struct {
		GinMode   string `env:"GIN_MODE" env-default:"debug"`
		Domain    string `env:"DOMAIN" env-default:"localhost"`
		AdminUser struct {
			Username string `env:"ADMIN_USERNAME" env-default:"admin"`
			Password string `env:"ADMIN_PWD" env-default:"admin"`
		}
		Auth struct {
			SessionName string `env:"SESSION_NAME" env-required:"true"`
		}
		MaxLimitPage int `env:"MAX_LIMIT_PAGE" env-default:"50"`
		JWTToken     struct {
			JwtAccessKey          string `env:"JWT_ACCESS_KEY" env-required:"true"`
			JwtRefreshKey         string `env:"JWT_REFRESH_KEY" env-required:"true"`
			AccessTokenExpiresIn  int    `env:"ACCESS_TOKEN_EXPIRED_IN" env-required:"true"`
			RefreshTokenExpiresIn int    `env:"REFRESH_TOKEN_EXPIRED_IN" env-required:"true"`
			AccessTokenMaxAge     int    `env:"ACCESS_TOKEN_MAXAGE" env-required:"true"`
			RefreshTokenMaxAge    int    `env:"REFRESH_TOKEN_MAXAGE" env-required:"true"`
			MaxTokenKeys          int    `env:"MAX_TOKEN_KEYS" env-default:"5"`
		}
	}
	PostgreSQL struct {
		Username string `env:"POSTGRES_USER" env-required:"true"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		Host     string `env:"POSTGRES_HOST" env-required:"true"`
		Port     string `env:"POSTGRES_PORT" env-required:"true"`
		Database string `env:"POSTGRES_DB" env-required:"true"`
	}
	Redis struct {
		Host   string `env:"REDIS_HOST" env-required:"true"`
		Port   string `env:"REDIS_PORT" env-required:"true"`
		Secret string `env:"REDIS_SECRET" env-required:"true"`
		Size   int    `env:"REDIS_SIZE" env-required:"true"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "Go-rshok todo system"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
