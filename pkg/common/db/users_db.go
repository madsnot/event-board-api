package db

import (
	"example/event-board/pkg/common/models"
	"fmt"
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByEmail(ctx *gin.Context, dbPool *pgxpool.Pool, email string) (id int, pass string, _ error) {
	sqlQuery, _, _ := sq.Select("id", "user_password").From("users").Where(fmt.Sprintf("email = '%s'", email)).ToSql()
	queryRow := dbPool.QueryRow(ctx, sqlQuery)
	errScanQuery := queryRow.Scan(&id, &pass)
	return id, pass, errScanQuery
}

func CreateUser(ctx *gin.Context, dbPool *pgxpool.Pool, user *models.UsersList) {
	sqlQuery := fmt.Sprintf("INSERT INTO users (user_password,email,user_name,user_surname,"+
		"user_sex,birthday_date) VALUES ('%s','%s','%s','%s','%s','%s')", user.Password, user.Email,
		user.Name, user.Surname, user.Sex, user.BirthdayDate)
	_, errCreateUser := dbPool.Query(ctx, sqlQuery)
	if errCreateUser != nil {
		log.Print("errCreateUser: ", errCreateUser)
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return
	}
}
