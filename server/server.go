package server

import (
	"example/event-board/server/pkg/common/db"
	"example/event-board/server/pkg/controllers/users"
	"example/event-board/server/pkg/email"
	"example/event-board/server/pkg/hash"
	"example/event-board/server/pkg/routes"
	"example/event-board/server/pkg/tokens"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server struct {
	dbPort          string
	dbUrl           string
	hashAddition    string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	tokenSigningKey string
	emailAddr       string
	emailPass       string
	emailHost       string
	emailPort       string
}

func InitServer() *Server {
	viper.SetConfigFile("./server/pkg/common/envs/.env")
	viper.ReadInConfig()
	ttl := viper.Get("ACCESS_TOKEN_TTL").(string)
	accessTokenTTL, _ := strconv.Atoi(ttl)
	ttl = viper.Get("REFRESH_TOKEN_TTL").(string)
	refreshTokenTTL, _ := strconv.Atoi(ttl)
	return &Server{
		dbPort:          viper.Get("PORT").(string),
		dbUrl:           viper.Get("DB_URL").(string),
		hashAddition:    viper.Get("HASH_ADDITION").(string),
		accessTokenTTL:  time.Hour * time.Duration(accessTokenTTL),
		refreshTokenTTL: time.Hour * time.Duration(refreshTokenTTL),
		tokenSigningKey: viper.Get("SIGNING_KEY").(string),
		emailAddr:       viper.Get("EMAIL_ADDR").(string),
		emailPass:       viper.Get("EMAIL_PASS").(string),
		emailHost:       viper.Get("EMAIL_HOST").(string),
		emailPort:       viper.Get("EMAIL_PORT").(string),
	}
}

func (server *Server) Run() {
	route := gin.Default()
	dbPool, errDBInit := db.Init(server.dbUrl)
	if errDBInit != nil {
		log.Fatal(errDBInit)
	}
	defer dbPool.Close()
	hasher := hash.NewHasher(server.hashAddition)
	token := tokens.NewTokenInfo(server.accessTokenTTL, server.refreshTokenTTL, server.tokenSigningKey)
	emailService := email.NewEmailService(server.emailAddr, server.emailPass, server.emailHost, server.emailPort)
	userService := users.NewUserService(hasher, token, emailService)
	routes.RegisterRoutes(route, dbPool, userService)

	route.Run(server.dbPort)
}
