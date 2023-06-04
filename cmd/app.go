package main

import (
	"log"

	"github.com/madsnot/event-board-api/internal/config"
	db "github.com/madsnot/event-board-api/internal/repository"
	"github.com/madsnot/event-board-api/internal/routes"
	"github.com/madsnot/event-board-api/internal/transport/users"
	"github.com/madsnot/event-board-api/pkg/email"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/tokens"

	"github.com/gin-gonic/gin"
)

// type Server struct {
// 	config.DBConfig
// 	config.
// 	dbPort          string
// 	dbUrl           string
// 	hashSalt        string
// 	accessTokenTTL  time.Duration
// 	refreshTokenTTL time.Duration
// 	tokenSigningKey string
// 	emailAddr       string
// 	emailPass       string
// 	emailHost       string
// 	emailPort       string
// }

// func InitServer() *Server {
// 	viper.SetConfigFile("./server/pkg/common/envs/.env")
// 	viper.ReadInConfig()
// 	ttl := viper.Get("ACCESS_TOKEN_TTL").(string)
// 	accessTokenTTL, _ := strconv.Atoi(ttl)
// 	ttl = viper.Get("REFRESH_TOKEN_TTL").(string)
// 	refreshTokenTTL, _ := strconv.Atoi(ttl)
// 	return &Server{
// 		dbPort:          viper.Get("PORT").(string),
// 		dbUrl:           viper.Get("DB_URL").(string),
// 		hashSalt:        viper.Get("HASH_SALT").(string),
// 		accessTokenTTL:  time.Hour * time.Duration(accessTokenTTL),
// 		refreshTokenTTL: time.Hour * time.Duration(refreshTokenTTL),
// 		tokenSigningKey: viper.Get("SIGNING_KEY").(string),
// 		emailAddr:       viper.Get("EMAIL_ADDR").(string),
// 		emailPass:       viper.Get("EMAIL_PASS").(string),
// 		emailHost:       viper.Get("EMAIL_HOST").(string),
// 		emailPort:       viper.Get("EMAIL_PORT").(string),
// 	}
// }

func Run() {
	route := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbPool, errDBInit := db.Init(cfg.DataBaseCfg.DBUrl)
	if errDBInit != nil {
		log.Fatal(errDBInit)
	}

	defer dbPool.Close()

	hasher := hash.NewHasher(cfg.HashCfg)

	tokenizer := tokens.NewTokenizer(cfg.TokenCfg)

	emailService := email.NewEmailService(cfg.EmailCfg)

	userService := users.NewUserService(hasher, tokenizer, emailService)

	routes.RegisterRoutes(route, dbPool, userService)

	route.Run(cfg.DataBaseCfg.Port)
}
