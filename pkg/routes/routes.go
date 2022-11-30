package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type handlerPool struct {
	DataBasePool *pgxpool.Pool
}

func RegisterRoutes(routes *gin.Engine, dataBasePool *pgxpool.Pool) {
	// handler := &handlerPool{
	// 	DataBasePool: dataBasePool,
	// }

	routes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello World!")
	})

	// //User
	// route := routes.Group("user")
	// route.GET("/", handler.GetUserPage)
	// route.POST("/login", handler.Login)
	// route.POST("/create", handler.CreateUser)

	// //Events
	// route = routes.Group("events")
	// route.POST("/", handler.GetEvents)
	// route.GET("/:id", handler.GetEvent)
	// route.POST("/create", handler.CreateEvent)

}
