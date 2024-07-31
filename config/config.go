package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// You can retrieve any configuration value like this
	fmt.Println("Cipher key:", viper.GetString("cipher.key"))
}
