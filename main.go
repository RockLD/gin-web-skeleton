package main

import (
	"errors"
	"fmt"
	"gin-web-skeleton/app/config"
	"gin-web-skeleton/middleware"
	"gin-web-skeleton/model"
	"gin-web-skeleton/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var cfg = pflag.StringP("config", "c", "", "path")

func main() {
	g := gin.New()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	router.InitRouter(g, middleware.Cors())

	go func() {
		if err := pingServer(); err != nil {
			fmt.Println("The router bas been deployed successfully")
		}
	}()

	model.DB.Init()
	defer model.DB.Close()

	fmt.Println(viper.GetString("jwt_secret"))

	fmt.Println(viper.Get("db"))
	g.Run(":8090")
}

func pingServer() error {
	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:8090" + "/check/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
