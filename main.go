package main

import (
	"gin-web-skeleton/router"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	g := gin.New()
	router.InitRouter(g)

	log.Println(http.ListenAndServe(":8090", g))
}
