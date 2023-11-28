package adapters

import (
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/transport/http/dto"
)

func AdaptSessionBmodelToSignInResponse(session models.Session) dto.SignInResponse {
	return dto.SignInResponse{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
	}
}
