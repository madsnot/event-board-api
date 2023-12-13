package v1

import (
	"encoding/json"
	"github.com/madsnot/event-board-api/internal/transport/http/dto"
	"net/http"

	"github.com/madsnot/event-board-api/internal/transport/http/adapters"
)

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequestToSignInRequestDTO(r)
	if err != nil {
		writeError(w, err)
	}

	ctx := r.Context()

	user, err := h.usecase.User.GetUserByEmail(ctx, req.Email)
	if err != nil {
		writeError(w, err)
	}

	hashPass, err := h.hasher.Hash(req.Password)
	if err != nil {
		writeError(w, err)
	}

	if !validPassword(user.Password, hashPass) {
		writeError(w, err)
	}

	session, err := h.usecase.Auth.CreateSession(r.Context(), user.ID)
	if err != nil {
		writeError(w, err)
	}

	respDTO := adapters.AdaptSessionBmodelToSignInResponse(session)

	respBody, err := json.Marshal(respDTO)
	if err != nil {
		writeError(w, err)
	}

	if _, err = w.Write(respBody); err != nil {
		writeError(w, err)
	}

	writeOK(w, dto.SignInResponse{})
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
