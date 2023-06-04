package users

import (
	"log"
	"net/http"

	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/domain/usecase"
	db "github.com/madsnot/event-board-api/internal/repository"
	"github.com/madsnot/event-board-api/pkg/email"
	"github.com/madsnot/event-board-api/pkg/hash"
	"github.com/madsnot/event-board-api/pkg/tokens"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	usecase *usecase.UserUsecase
	hasher  *hash.Hasher
}

func NewUserService(hasher *hash.Hasher, tokenizer *tokens.Tokenizer, email *email.Email) *UserService {
	return &UserService{
		usecase: usecase.NewUserUsecase(hasher, tokenizer, email),
		hasher:  hasher,
	}
}

func (us *UserService) SignIn(ctx *gin.Context, dbPool *pgxpool.Pool) {
	var (
		user              models.UsersList
		errGetUserByEmail error
	)

	errBindJSON := ctx.BindJSON(&user)
	if errBindJSON != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return
	}

	passHash, passErr := us.hasher.Hash(user.Password)
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
	token, errToken := us.usecase.CreateSession(ctx, user.ID, dbPool)
	if errToken != nil {
		log.Println("errToken: ", errToken)
		ctx.JSON(http.StatusInternalServerError, errToken)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
