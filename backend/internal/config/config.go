package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	MYSQL struct {
		DSN string `json:"DSN"`
	}
	Swagger struct {
		Host string `json:"Host"`
	}
}
