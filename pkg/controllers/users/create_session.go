package users

import (
	"example/event-board/pkg/common/db"
	"example/event-board/pkg/tokens"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (userService *UserService) createSession(ctx *gin.Context, userId int,
	dbPool *pgxpool.Pool) (_ *tokens.Token, err error) {

	var token tokens.Token
	tokenTemp := fmt.Sprintf("EVENTBOARDUSER%d", userId)
	token.AccessToken, err = userService.tokenInfo.NewAccessToken(tokenTemp)
	if err != nil {
		return &token, err
	}
	token.RefreshToken, err = userService.tokenInfo.NewRefreshToken()
	if err != nil {
		return &token, err
	}
	err = db.CreateSession(ctx, dbPool, userId, token.RefreshToken)
	if err != nil {
		return &token, err
	}
	return &token, nil
}
