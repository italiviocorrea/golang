package config

import (
	"github.com/spf13/viper"
	"log"
)

type Settings struct {
	DbHost         string `mapstrucute:"DBHOST"`
	DbPwd          string `mapstrucute:"DBPWD"`
	DbUser         string `mapstrucute:"DBUSER"`
	DbName         string `mapstrucute:"DBNAME"`
	Env            string `mapstrucute:"ENV"`
	SrvHost        string `mapstrucute:"SRVHOST"`
	SrvPort        string `mapstrucute:"SRVPORT"`
	JaegerEndpoint string `mapstrucute:"JAEGERENDPOINT"`
}

func New() *Settings {

	var cfg Settings

	viper.SetConfigFile("app.env")
	viper.SetConfigType("env")

	//viper.SetEnvPrefix("SDT")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("Não encontrado arquivo env, usando variáveis de ambiente", err)
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		log.Fatalf("Erro ao tentar ler configurações")
	}

	return &cfg
}
