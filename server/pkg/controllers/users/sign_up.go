package users

import (
	"example/event-board/server/pkg/common/db"
	"example/event-board/server/pkg/common/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (userService *UserService) SignUp(ctx *gin.Context, dbPool *pgxpool.Pool) {
	var (
		user              models.UsersList
		errGetUserByEmail error
	)
	errBindJSON := ctx.BindJSON(&user)
	if errBindJSON != nil {
		log.Print("errBindJSON: ", errBindJSON)
		ctx.JSON(http.StatusInternalServerError, errBindJSON)
		return
	}
	passHash, passErr := userService.hasher.Hash(user.Password)
	if passErr != nil {
		log.Print(passErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return
	}
	user.ID, user.Password, errGetUserByEmail = db.GetUserByEmail(ctx, dbPool, user.Email)
	if errGetUserByEmail == nil {
		log.Print("errGetUserByEmail:", errGetUserByEmail)
		ctx.JSON(http.StatusNotFound, gin.H{"response: ": "This user already exists"})
		return
	}
	user.Password = passHash
	db.CreateUser(ctx, dbPool, &user)
	db.CreateStudent(ctx, dbPool, &user.StudentInfo)
	userService.email.Recipient = user.Email
	status := userService.email.Verify(ctx)
	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{"response: ": "Мalidation code is incorrect"})
		return
	}
}
