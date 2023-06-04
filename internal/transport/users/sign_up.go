package users

import (
	"log"
	"net/http"

	"github.com/madsnot/event-board-api/internal/domain/models"
	db "github.com/madsnot/event-board-api/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (us *UserService) SignUp(ctx *gin.Context, dbPool *pgxpool.Pool) {
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

	passHash, passErr := us.hasher.Hash(user.Password)
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

	// userService.email.Recipient = user.Email

	// status := us.email.Verify(ctx)
	// if !status {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"response: ": "Мalidation code is incorrect"})
	// 	return
	// }
}
