package util

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type (
	Configuration struct {
		Server   string `default:"0.0.0.0:7878"`
		DBHost   string `default:"s1602.ms"`
		DBPort   int    `default:"9042"`
		DBUser   string `default:"nf3e"`
		DBPwd    string `default:"senha#123"`
		Database string `default:"nf3e"`
		LogLevel int    `default:"4"`
	}
)

// AppConfig holds the configuration values from config.json file
var AppConfig Configuration

// Initialize AppConfig
func initConfig(varEnvPrefix string) {
	loadAppEnv(varEnvPrefix)
}

// Reads configuration da env vars.
func loadAppEnv(varEnvPrefix string) {

	AppConfig = Configuration{}
	err := envconfig.Process(varEnvPrefix, &AppConfig)
	if err != nil {
		log.Fatalf("[loadAppEnv]: %s\n", err)
	}
	log.Println(AppConfig)
}
