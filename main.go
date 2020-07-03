package main

import (
	"errors"
	"fmt"
	"gin-web-skeleton/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	g := gin.New()

	router.InitRouter(g)

	go func() {
		if err := pingServer(); err != nil {
			fmt.Println("The router bas been deployed successfully")
		}
	}()

	fmt.Println(viper.GetString("addr"))
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
