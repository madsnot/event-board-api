package users

import (
	"example/event-board/pkg/common/models"
	"fmt"
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (userService *UserService) SignIn(ctx *gin.Context, dbPool *pgxpool.Pool) {
	var (
		user   models.UsersList
		dbUser models.UsersList
	)
	errBindJSON := ctx.BindJSON(&user)
	if errBindJSON != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return
	}
	sqlQuery, _, _ := sq.Select("id", "user_password").From("users").Where(fmt.Sprintf("login_name = '%s'", user.Login)).ToSql()
	queryRow := dbPool.QueryRow(ctx, sqlQuery)
	if errScanQuery := queryRow.Scan(&dbUser.ID, &dbUser.Password); errScanQuery != nil {
		log.Print("errScanQuery:", errScanQuery)
		ctx.JSON(http.StatusNotFound, gin.H{"response: ": "User not found"})
		return
	}
	passHash, passErr := userService.hasher.Hash(user.Password)
	if passErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return
	}
	if passHash != dbUser.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{"response: ": "Wrong password!"})
		return
	}
	token, errToken := userService.createSession(ctx, dbUser.ID, dbPool)
	if errToken != nil {
		log.Println("errToken: ", errToken)
		ctx.JSON(http.StatusInternalServerError, errToken)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
