package usecase

import (
	"fmt"

	"github.com/madsnot/event-board-api/internal/domain/models"
	db "github.com/madsnot/event-board-api/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (uc *UserUsecase) CreateSession(ctx *gin.Context, userId int, dbPool *pgxpool.Pool) (_ *models.Token, err error) {
	var token models.Token

	tokenTemp := fmt.Sprintf("EVENTBOARDUSER%d", userId)

	token.AccessToken, err = uc.tokenInfo.NewAccessToken(tokenTemp)
	if err != nil {
		return &token, err
	}

	token.RefreshToken, err = uc.tokenInfo.NewRefreshToken()
	if err != nil {
		return &token, err
	}

	err = db.CreateSession(ctx, dbPool, userId, token.RefreshToken)
	if err != nil {
		return &token, err
	}

	return &token, nil
}
