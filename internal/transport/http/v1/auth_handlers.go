package v1

import (
	"encoding/json"
	"net/http"

	"github.com/madsnot/event-board-api/internal/transport/http/adapters"
)

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequestToSignInRequestDTO(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	ctx := r.Context()

	user, err := h.usecase.User.GetUserByEmail(ctx, req.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	hashPass, err := h.hasher.Hash(req.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if !validPassword(user.Password, hashPass) {
		w.WriteHeader(http.StatusBadRequest)
	}

	session, err := h.usecase.Auth.CreateSession(r.Context(), user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	respDTO := adapters.AdaptSessionBmodelToSignInResponse(session)

	respBody, err := json.Marshal(respDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if _, err = w.Write(respBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	//reqDTO, err := parseRequestToRefreshTokenDTO(r)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//}
	//
	//session, err := h.usecase.Auth.GetSession(r.Context())
}

func (h *handler) Logout(w http.ResponseWriter, r *http.Request) {

}
