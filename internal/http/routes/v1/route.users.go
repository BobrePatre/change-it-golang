package v1

import (
	V1Usecase "change-it/internal/business/usecases/v1"
	V1PostgresRepository "change-it/internal/datasources/repositories/postgres/v1"
	V1Handlers "change-it/internal/http/handlers/v1"
	"change-it/internal/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UsersRoutes struct {
	V1Handler V1Handlers.UsersHandler
	router    *gin.RouterGroup
	db        *sqlx.DB
}

func NewUserRoutes(router *gin.RouterGroup, db *sqlx.DB) *UsersRoutes {
	V1UserRepository := V1PostgresRepository.NewUserRepository(db)
	V1UsersUsecase := V1Usecase.NewUserUsecase(V1UserRepository)
	V1UserHandler := V1Handlers.NewUsersHandler(V1UsersUsecase)

	return &UsersRoutes{V1Handler: V1UserHandler, router: router, db: db}
}

func (r *UsersRoutes) RegisterRoutes() {
	V1PetitionRoute := r.router.Group("/user")
	{
		V1PetitionRoute.GET("/likes", middlewares.KycloakAuthMiddleware(), r.V1Handler.GetLikedPetitions)
		V1PetitionRoute.GET("/voices", middlewares.KycloakAuthMiddleware(), r.V1Handler.GetVoicedPetitions)
		V1PetitionRoute.GET("", middlewares.KycloakAuthMiddleware(), r.V1Handler.GetOwnedPetitions)
	}

}
