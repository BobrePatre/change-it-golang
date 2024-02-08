package v1

import (
	V1Usecase "change-it/internal/business/usecases/v1"
	V1PostgresRepository "change-it/internal/datasources/repositories/postgres/v1"
	V1Handlers "change-it/internal/http/handlers/v1"
	"change-it/internal/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type UsersRoutes struct {
	V1Handler      V1Handlers.UsersHandler
	router         *gin.RouterGroup
	db             *sqlx.DB
	AuthMiddleware func(roles ...string) gin.HandlerFunc
}

func NewUserRoutes(router *gin.RouterGroup, db *sqlx.DB, redis *redis.Client) *UsersRoutes {
	V1UserRepository := V1PostgresRepository.NewUserRepository(db)
	V1UsersUsecase := V1Usecase.NewUserUsecase(V1UserRepository)
	V1UserHandler := V1Handlers.NewUsersHandler(V1UsersUsecase)
	KeycloakAuthMiddleware := middlewares.KeycloakAuthMiddleware(redis)

	return &UsersRoutes{V1Handler: V1UserHandler, router: router, db: db, AuthMiddleware: KeycloakAuthMiddleware}
}

func (r *UsersRoutes) RegisterRoutes() {
	V1PetitionRoute := r.router.Group("/user")
	{
		V1PetitionRoute.GET("/likes", r.AuthMiddleware(), middlewares.Pagination(), r.V1Handler.GetLikedPetitions)
		V1PetitionRoute.GET("/voices", r.AuthMiddleware(), middlewares.Pagination(), r.V1Handler.GetVoicedPetitions)
		V1PetitionRoute.GET("", r.AuthMiddleware(), middlewares.Pagination(), r.V1Handler.GetOwnedPetitions)
	}

}
