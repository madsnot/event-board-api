package v1

import (
	"github.com/gorilla/mux"
	"github.com/madsnot/event-board-api/internal/config"
	"github.com/madsnot/event-board-api/internal/domain/usecase"
	"github.com/madsnot/event-board-api/pkg/hash"
	"net/http"
)

type handler struct {
	usecase usecase.Usecase
	cfg     config.Config
	hasher  *hash.Hasher
}

func NewHandler(usecase usecase.Usecase, cfg config.Config) *handler {
	return &handler{
		usecase: usecase,
		cfg:     cfg,
		hasher:  hash.NewHasher(cfg.HashCfg),
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/auth/signUp", h.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/auth/signIn", h.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/auth/refresh", h.RefreshToken).Methods(http.MethodGet)
	router.HandleFunc("/auth/logout", h.Logout).Methods(http.MethodGet)

	router.HandleFunc("/api/events", h.GetEventList).Methods(http.MethodGet)
	router.HandleFunc("/api/events/create", h.CreateEvent).Methods(http.MethodPost)
	router.HandleFunc("/api/events/id", h.GetEvent).Methods(http.MethodGet)
	router.HandleFunc("/api/events/id/update", h.UpdateEvent).Methods(http.MethodPost)
	router.HandleFunc("/api/events/id/close", h.CloseEvent).Methods(http.MethodGet)
}
