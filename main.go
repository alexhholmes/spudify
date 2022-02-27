package main

import (
	"log"

	"example/spudify/app"
	"example/spudify/controllers"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Failed to read config file: %s", err)
		log.Panicf("helloworld");

	}
}

func main() {
	a := app.App{}
	a.Initialize(viper.GetString("DB.address"),
		viper.GetString("DB.port"),
		viper.GetString("DB.username"),
		viper.GetString("DB.password"),
		viper.GetString("DB.db_name"))
	controllers.InitDB(a.DB)

	a.Run(viper.GetString("Server.address"),
		viper.GetString("Server.port"))
}
