package component

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config ConfigT

type ConfigT struct {
	This      this      `mapstructure:"this"`
	MyRequest myRequest `mapstructure:"my_request"`
	RsaConfig rsaConfig `mapstructure:"rsa_config"`
}

type rsaConfig struct {
	PrivateKey string `mapstructure:"private_key"`
	PublicKey  string `mapstructure:"public_key"`
}

type myRequest struct {
	Url  string `mapstructure:"url"`
	Port string `mapstructure:"port"`
}

type this struct {
	Port          string `mapstructure:"port"`
	CaPemPath     string `mapstructure:"ca_pem_path"`
	ClientPemPath string `mapstructure:"client_pem_path"`
	ClientKeyPath string `mapstructure:"client_key_path"`
}

func LoadConfig(configPath string) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading model file: %v", err))
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Sprintf("Error unmarshaling model: %v", err))
	}
}
