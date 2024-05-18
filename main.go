package main

import (
	"github.com/DangerZombie/case-study-dealls/initialization"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()

	// init DB Connection
	db, err := initialization.DbInit()
	if err != nil {
		panic(err)
	}

	initialization.ServerInit(db)
}
