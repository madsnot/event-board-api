package main

import (
	"example/event-board/pkg/common/db"
	"example/event-board/pkg/routes"
	"log"

	//"fmt"

	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
}

func main() {
	route := gin.Default()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	dbPool, errDBInit := db.Init(dbUrl)
	if errDBInit != nil {
		log.Fatal(errDBInit)
	}
	defer dbPool.Close()
	routes.RegisterRoutes(route, dbPool)

	route.Run(":" + port)
}
