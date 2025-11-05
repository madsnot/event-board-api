package v1

import (
	"github.com/gorilla/mux"
	"github.com/madsnot/event-board-api/internal/config"
	"github.com/madsnot/event-board-api/internal/domain/usecase"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/token"
	"net/http"
)

type handler struct {
	usecase   usecase.Usecase
	cfg       config.Config
	hasher    *hash.Hasher
	tokenizer *token.Tokenizer
}

func NewHandler(usecase usecase.Usecase, cfg config.Config) *handler {
	tokenizer, _ := token.NewTokenizer(cfg.TokenCfg)

	return &handler{
		usecase:   usecase,
		cfg:       cfg,
		hasher:    hash.NewHasher(cfg.HashCfg),
		tokenizer: tokenizer,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/auth/signUp", h.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/auth/signIn", h.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/auth/refresh", AuthMiddleware(h.RefreshToken, h.tokenizer)).Methods(http.MethodGet)
	router.HandleFunc("/auth/logout", AuthMiddleware(h.Logout, h.tokenizer)).Methods(http.MethodGet)

	router.HandleFunc("/api/events", h.GetEventList).Methods(http.MethodGet)
	router.HandleFunc("/api/events/create", h.CreateEvent).Methods(http.MethodPost)
	router.HandleFunc("/api/events/id", AuthMiddleware(h.GetEvent, h.tokenizer)).Methods(http.MethodGet)
	router.HandleFunc("/api/events/id/update", AuthMiddleware(h.UpdateEvent, h.tokenizer)).Methods(http.MethodPost)
	router.HandleFunc("/api/events/id/close", AuthMiddleware(h.CloseEvent, h.tokenizer)).Methods(http.MethodGet)
}
