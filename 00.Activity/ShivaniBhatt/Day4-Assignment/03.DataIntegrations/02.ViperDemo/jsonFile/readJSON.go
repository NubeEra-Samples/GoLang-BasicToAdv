package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	port := viper.Get("prod.port")

	fmt.Println(port)
}
