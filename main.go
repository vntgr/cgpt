package main

import (
	"github.com/spf13/viper"

	"git.mysticmode.net/mysticmode/kily/cmd"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("kily")
	viper.SetConfigType("yml")
	viper.SetConfigFile("./kily.yml")
	viper.ReadInConfig()

	cmd.Execute()
}
