package users

import (
	"example/event-board/pkg/tokens"
	"fmt"
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (userService *UserService) createSession(ctx *gin.Context, userId int,
	dbPool *pgxpool.Pool) (token *tokens.Token, err error) {
	tokenTemp := fmt.Sprintf("EVENTBOARDUSER%d", userId)
	token.AccessToken, err = userService.tokenInfo.NewAccessToken(tokenTemp)
	if err != nil {
		return nil, err
	}
	token.RefreshToken, err = userService.tokenInfo.NewRefreshToken()
	sqlQuery, _, _ := sq.Insert("users_sessions").Columns("user_id", "refresh_token").Values(userId, token.RefreshToken).ToSql()
	_, errQuery := dbPool.Query(ctx, sqlQuery)
	if errQuery != nil {
		log.Print("errQuery:", errQuery)
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return
	}
	return token, nil
}
