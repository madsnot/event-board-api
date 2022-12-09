package users

import (
	"example/event-board/pkg/tokens"
	"fmt"
	"log"

	//sq "github.com/Masterminds/squirrel"
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
	sqlQuery := fmt.Sprintf("INSERT INTO users_sessions (user_id,refresh_token) VALUES (%d,'%x')", userId, token.RefreshToken)
	//sqlQuery, _, _ := sq.Insert("users_sessions").Columns("user_id", "refresh_token").Values(userId, token.RefreshToken).ToSql()
	_, errQuery := dbPool.Query(ctx, sqlQuery)
	if errQuery != nil {
		log.Print("errQuery: ")
		return &token, errQuery
	}
	return &token, nil
}
