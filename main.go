package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-prac-site/internal/router"
	"net/http"
)

func main() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	err := http.ListenAndServe(fmt.Sprintf(":%v", viper.Get("http.port")), router.NewNegroni())
	if err != nil {
		panic(err)
	}
}


