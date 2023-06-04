package db

import (
	"fmt"
	"log"
	"net/http"

	"github.com/madsnot/event-board-api/internal/domain/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateStudent(ctx *gin.Context, dbPool *pgxpool.Pool, student *models.Student) {
	sqlQuery := fmt.Sprintf("INSERT INTO students (school,group_num,dorm)"+
		" VALUES ('%s','%s','%s')", student.School, student.Group, student.Dorm)
	_, errCreateStudent := dbPool.Query(ctx, sqlQuery)
	if errCreateStudent != nil {
		log.Print("errCreateStudent: ", errCreateStudent)
		ctx.JSON(http.StatusInternalServerError, gin.H{"response: ": "Internal server error"})
		return
	}
}
