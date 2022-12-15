package routes

import (
	"example/event-board/server/pkg/controllers/users"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type handlerPool struct {
	DataBasePool *pgxpool.Pool
}

func (handler *handlerPool) validPool(controller func(*gin.Context, *pgxpool.Pool)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		controller(ctx, handler.DataBasePool)
	}
}

func RegisterRoutes(routes *gin.Engine, dataBasePool *pgxpool.Pool, service *users.UserService) {
	handler := &handlerPool{
		DataBasePool: dataBasePool,
	}

	routes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello World!")
	})

	//User
	route := routes.Group("/user")
	//	route.GET("/", users.GetUserPage)
	route.POST("/signIn", handler.validPool(service.SignIn))
	route.POST("/signUp", handler.validPool(service.SignUp))

	// //Events
	// route = routes.Group("/events")
	// route.POST("/", handler.GetEvents)
	// route.GET("/:id", handler.GetEvent)
	// route.POST("/create", handler.CreateEvent)

}
