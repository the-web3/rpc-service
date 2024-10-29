package flags

import "github.com/urfave/cli/v2"

const evnVarPrefix = "SIGNATURE"

func prefixEnvVars(name string) []string {
	return []string{evnVarPrefix + "_" + name}
}

var (
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path for database migrations",
		EnvVars: prefixEnvVars("MIGRATIONS_DIR"),
	}
	// RpcHostFlag RPC Service
	RpcHostFlag = &cli.StringFlag{
		Name:     "rpc-host",
		Usage:    "The host of the rpc",
		EnvVars:  prefixEnvVars("RPC_HOST"),
		Required: true,
	}
	RpcPortFlag = &cli.IntFlag{
		Name:     "rpc-port",
		Usage:    "The port of the rpc",
		EnvVars:  prefixEnvVars("RPC_PORT"),
		Value:    8987,
		Required: true,
	}
	// MetricsHostFlag Metrics
	MetricsHostFlag = &cli.StringFlag{
		Name:     "metrics-host",
		Usage:    "The host of the metrics",
		EnvVars:  prefixEnvVars("METRICS_HOST"),
		Required: true,
	}
	MetricsPortFlag = &cli.IntFlag{
		Name:     "metrics-port",
		Usage:    "The port of the metrics",
		EnvVars:  prefixEnvVars("METRICS_PORT"),
		Value:    7214,
		Required: true,
	}
	HttpHostFlag = &cli.StringFlag{
		Name:     "http-host",
		Usage:    "The host of the rest api",
		EnvVars:  prefixEnvVars("HTTP_HOST"),
		Required: true,
	}
	HttpPortFlag = &cli.IntFlag{
		Name:     "http-port",
		Usage:    "The port of the rest api",
		EnvVars:  prefixEnvVars("HTTP_PORT"),
		Required: true,
	}

	// DbHostFlag Database
	DbHostFlag = &cli.StringFlag{
		Name:     "master-db-host",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("DB_HOST"),
		Required: true,
	}
	DbPortFlag = &cli.IntFlag{
		Name:     "master-db-port",
		Usage:    "The port of the master database",
		EnvVars:  prefixEnvVars("DB_PORT"),
		Required: true,
	}
	DbUserFlag = &cli.StringFlag{
		Name:     "master-db-user",
		Usage:    "The user of the master database",
		EnvVars:  prefixEnvVars("DB_USER"),
		Required: true,
	}
	DbPasswordFlag = &cli.StringFlag{
		Name:     "master-db-password",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("DB_PASSWORD"),
		Required: true,
	}
	DbNameFlag = &cli.StringFlag{
		Name:     "master-db-name",
		Usage:    "The db name of the master database",
		EnvVars:  prefixEnvVars("DB_NAME"),
		Required: true,
	}
)

var requireFlags = []cli.Flag{
	MigrationsFlag,
	RpcHostFlag,
	RpcPortFlag,
	MetricsPortFlag,
	MetricsHostFlag,
	HttpHostFlag,
	HttpPortFlag,
	DbHostFlag,
	DbPortFlag,
	DbUserFlag,
	DbPasswordFlag,
	DbNameFlag,
}

var optionalFlags = []cli.Flag{}

func init() {
	Flags = append(requireFlags, optionalFlags...)
}

var Flags []cli.Flag
