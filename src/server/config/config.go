package config

import (
	"template/core/log"

	goenv "github.com/evilwire/go-env"
)

type configServer struct {
	HttpPort  string `env:"HTTP_PORT"`
	HttpsPort string `env:"HTTPS_PORT"`
}

type configDatabaseServer struct {
	DbUsername               string `env:"USERNAME"`
	DbPassword               string `env:"PASSWORD"`
	DbAddress                string `env:"ADDRESS"`
	DbPort                   string `env:"PORT"`
	DbSchema                 string `env:"SCHEMA"`
	DbName                   string `env:"DBNAME"`
	DbMaxOpenConnections     int    `env:"MAX_OPEN_CONNS"`
	DbMaxIdleConnections     int    `env:"MAX_IDLE_CONNS"`
	DbMaxConnectionsLifetime int    `env:"MAX_CONNS_LIFETIME"`
	DbMaxConnectionsIdleTime int    `env:"MAX_CONNS_IDLE_TIME"`
	DbSilent                 bool   `env:"SILENT"`
}

type Config struct {
	Server          configServer         `env:"SERVER_"`
	Database        configDatabaseServer `env:"DATABASE_"`
	CertificatePath string               `env:"HTTPS_CERTIFICATE_PATH"`
	KeyPath         string               `env:"HTTPS_KEY_PATH"`
}

var cfg = Config{}

var configInitialized = false

func initConfig() (config *Config) {
	marshaller := goenv.DefaultEnvMarshaler{
		Environment: goenv.NewOsEnvReader(),
	}
	err := marshaller.Unmarshal(&cfg)
	// cfg.fillUndefinedWithDefaults()
	if err != nil {
		log.Errorf("[initConfig]: %s", err)
	} else {
		log.Info("[initConfig]: Configuration initiated successfuly")
	}
	return &cfg
}

func GetConfig() *Config {
	if !configInitialized {
		cfg := initConfig()
		configInitialized = true
		return cfg
	}
	return &cfg
}
