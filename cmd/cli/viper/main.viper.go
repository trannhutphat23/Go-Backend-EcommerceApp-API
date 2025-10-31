package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("../../../config/") // path to config
	viper.SetConfigName("local")            // file name
	viper.SetConfigType("yaml")             // extension

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	fmt.Println("Server port:", viper.GetInt("server.port"))
	fmt.Println("Jwt Key:", viper.GetString("security.jwt.key"))

	// configure structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuarion %v", err)
	}

	fmt.Println("Config port:", config.Server.Port)

	for _, db := range config.Database {
		fmt.Printf("Databases user: %s, password: %s, hostname: %s \n", db.User, db.Password, db.Host)
	}
}
