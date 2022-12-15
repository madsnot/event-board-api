package main

import (
	"example/event-board/pkg/common/db"
	"example/event-board/pkg/controllers/users"
	"example/event-board/pkg/email"
	"example/event-board/pkg/hash"
	"example/event-board/pkg/routes"
	"example/event-board/pkg/tokens"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()
}

func main() {
	route := gin.Default()

	dbPort := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	hashAddition := viper.Get("HASH_ADDITION").(string)
	ttl := viper.Get("ACCESS_TOKEN_TTL").(string)
	ttlInt, _ := strconv.Atoi(ttl)
	accessTokenTTL := time.Hour * time.Duration(ttlInt)
	ttl = viper.Get("REFRESH_TOKEN_TTL").(string)
	ttlInt, _ = strconv.Atoi(ttl)
	refreshTokenTTL := time.Hour * time.Duration(ttlInt)
	signingKey := viper.Get("SIGNING_KEY").(string)
	emailAddr := viper.Get("EMAIL_ADDR").(string)
	emailPass := viper.Get("EMAIL_PASS").(string)
	emailHost := viper.Get("EMAIL_HOST").(string)
	emailPort := viper.Get("EMAIL_PORT").(string)
	dbPool, errDBInit := db.Init(dbUrl)
	if errDBInit != nil {
		log.Fatal(errDBInit)
	}
	defer dbPool.Close()
	hasher := hash.NewHasher(hashAddition)
	token := tokens.NewTokenInfo(accessTokenTTL, refreshTokenTTL, signingKey)
	emailService := email.NewEmailService(emailAddr, emailPass, emailHost, emailPort)
	userService := users.NewUserService(hasher, token, emailService)
	routes.RegisterRoutes(route, dbPool, userService)

	route.Run(dbPort)
}
