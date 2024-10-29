package config

import (
	"github.com/urfave/cli/v2"

	"github.com/the-web3/rpc-service/flags"
)

type Config struct {
	Migrations    string
	Database      DBConfig
	RpcServer     ServerConfig
	HTTPServer    ServerConfig
	MetricsServer ServerConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type ServerConfig struct {
	Host string
	Port int
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		Migrations: ctx.String(flags.MigrationsFlag.Name),
		Database: DBConfig{
			Host:     ctx.String(flags.DbHostFlag.Name),
			Port:     ctx.Int(flags.DbPortFlag.Name),
			Name:     ctx.String(flags.DbNameFlag.Name),
			User:     ctx.String(flags.DbUserFlag.Name),
			Password: ctx.String(flags.DbPasswordFlag.Name),
		},
		RpcServer: ServerConfig{
			Host: ctx.String(flags.RpcHostFlag.Name),
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
		HTTPServer: ServerConfig{
			Host: ctx.String(flags.HttpHostFlag.Name),
			Port: ctx.Int(flags.HttpPortFlag.Name),
		},
		MetricsServer: ServerConfig{
			Host: ctx.String(flags.MetricsHostFlag.Name),
			Port: ctx.Int(flags.MetricsPortFlag.Name),
		},
	}
}
