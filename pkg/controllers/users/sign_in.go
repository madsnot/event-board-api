package users

import (
	"example/event-board/pkg/common/db"
	"example/event-board/pkg/common/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (userService *UserService) SignIn(ctx *gin.Context, dbPool *pgxpool.Pool) {
	var (
		user              models.UsersList
		errGetUserByEmail error
	)
	errBindJSON := ctx.BindJSON(&user)
	if errBindJSON != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return
	}
	passHash, passErr := userService.hasher.Hash(user.Password)
	if passErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return
	}
	user.ID, user.Password, errGetUserByEmail = db.GetUserByEmail(ctx, dbPool, user.Email)
	if errGetUserByEmail != nil {
		log.Print("errGetUserByEmail: ", errGetUserByEmail)
		ctx.JSON(http.StatusNotFound, gin.H{"response: ": "User not found"})
		return
	}
	if passHash != user.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{"response: ": "Wrong password!"})
		return
	}
	token, errToken := userService.createSession(ctx, user.ID, dbPool)
	if errToken != nil {
		log.Println("errToken: ", errToken)
		ctx.JSON(http.StatusInternalServerError, errToken)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
