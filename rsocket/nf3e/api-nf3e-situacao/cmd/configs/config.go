package configs

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/configs/logs"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"sync"
)

type (
	appConfig struct {
		// API Config
		LogLevel     int
		TpAmb        string
		CUF          string
		NSiteAutoriz int
		// ApiServer Config
		Server string
		Port   int
		// Database Config
		DBHost   string
		DBPort   int
		DBUser   string
		DBPwd    string
		Database string
		// Redis
		RedisHost     string
		RedisPort     string
		RedisDB       int
		RedisPassword string
	}
)

// config holds the configuration values from configs.json file
var config appConfig
var doOnce sync.Once

// Initialize config
func init() {
	doOnce.Do(func() {
		config = load()
	})
}

// Reads configuration da env vars.
func load() appConfig {
	var envVarsPrefix = defaultPrefixEnvVars()

	log.Println("Inicializa variaveis de ambiente, caso nao exista")
	setVarEnvs(envVarsPrefix)

	log.Println("Carregando as configuracoes da aplicacao.")
	config = appConfig{}

	err := envconfig.Process(envVarsPrefix+"", &config)

	if err != nil {
		log.Fatalf("[load]: %s\n", err)
	}

	log.Println("Configurando nivel de log...")
	logs.SetLogLevel(logs.Level(config.LogLevel))

	return config
}

func defaultPrefixEnvVars() string {
	envVarsPrefix := os.Getenv("API_PREFIX_ENV_VARS")

	if envVarsPrefix == "" {
		envVarsPrefix = "API"
	}
	return envVarsPrefix
}

func Get() appConfig {
	return config
}

func setVarEnvs(envVarsPrefix string) {

	if os.Getenv(envVarsPrefix+"_SERVER") == "" {
		os.Setenv(envVarsPrefix+"_SERVER", "localhost")
	}

	if os.Getenv(envVarsPrefix+"_PORT") == "" {
		os.Setenv(envVarsPrefix+"_PORT", "7878")
	}

	if os.Getenv(envVarsPrefix+"_DBHOST") == "" {
		os.Setenv(envVarsPrefix+"_DBHOST", "s1602.ms:9042")
	}

	if os.Getenv(envVarsPrefix+"_DBUSER") == "" {
		os.Setenv(envVarsPrefix+"_DBUSER", "admin")
	}

	if os.Getenv(envVarsPrefix+"_DBPWD") == "" {
		os.Setenv(envVarsPrefix+"_DBPWD ", "senha#123")
	}

	if os.Getenv(envVarsPrefix+"_DATABASE") == "" {
		os.Setenv(envVarsPrefix+"_DATABASE", "nf3e")
	}

	//if os.Getenv(envVarsPrefix+"_CONTEXT") == "" {
	//	os.Setenv(envVarsPrefix+"_CONTEXT ", "/nf3/v1")
	//}

	if os.Getenv(envVarsPrefix+"_LOGLEVEL") == "" {
		os.Setenv(envVarsPrefix+"_LOGLEVEL", "4")
	}

	if os.Getenv(envVarsPrefix+"_TPAMB") == "" {
		os.Setenv(envVarsPrefix+"_TPAMB", "2")
	}

	if os.Getenv(envVarsPrefix+"_CUF") == "" {
		os.Setenv(envVarsPrefix+"_CUF", "50")
	}

	if os.Getenv(envVarsPrefix+"_NSITEAUTORIZ") == "" {
		os.Setenv(envVarsPrefix+"_NSITEAUTORIZ", "0")
	}

}
