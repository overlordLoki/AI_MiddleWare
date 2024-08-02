package config

import (
	"github.com/spf13/viper" // Viper for configuration
	"fmt" // Format package
)

var Config struct {
    LLMAPIURL string `mapstructure:"llm_api_url"`
    ModelName string `mapstructure:"model_name"`
    Port      string `mapstructure:"port"`
}

func InitConfig() {
    viper.SetConfigName("config")  // name of config file (without extension)
    viper.SetConfigType("json")    // the format of the config file
    viper.AddConfigPath(".")       // look for config in the current directory

    if err := viper.ReadInConfig(); err != nil {
        panic(fmt.Errorf("fatal error reading config file: %w", err))
    }

    if err := viper.Unmarshal(&Config); err != nil {
        panic(fmt.Errorf("unable to decode into struct: %w", err))
    }
}