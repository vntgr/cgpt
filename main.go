package main

import (
	"git.mysticmode.net/mysticmode/kily/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("kily")
	viper.SetConfigType("yml")
	viper.SetConfigFile("./kily.yml")
	viper.ReadInConfig()

	cmd.Execute()
}
