package db

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateSession(ctx *gin.Context, dbPool *pgxpool.Pool, userId int, refreshToken string) error {
	sqlQuery := fmt.Sprintf("INSERT INTO users_sessions (user_id,refresh_token) VALUES (%d,'%x')", userId, refreshToken)
	//sqlQuery, _, _ := sq.Insert("users_sessions").Columns("user_id", "refresh_token").Values(userId, token.RefreshToken).ToSql()
	_, errCreateSession := dbPool.Query(ctx, sqlQuery)
	return errCreateSession
}
