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

type PetitionsRoutes struct {
	V1Handler      V1Handlers.PetitionsHandler
	router         *gin.RouterGroup
	db             *sqlx.DB
	AuthMiddleware func(roles ...string) gin.HandlerFunc
}

func NewPetitionRoute(router *gin.RouterGroup, db *sqlx.DB, redis *redis.Client) *PetitionsRoutes {
	V1PetitionRepository := V1PostgresRepository.NewPetitionRepository(db)
	V1UserRepository := V1PostgresRepository.NewUserRepository(db)
	V1PetitionUsecase := V1Usecase.NewPetitionUsecase(V1PetitionRepository, V1UserRepository)
	V1PetitionHandler := V1Handlers.NewPetitionsHandler(V1PetitionUsecase)
	KeycloakAuthMiddleware := middlewares.KeycloakAuthMiddleware(redis)

	return &PetitionsRoutes{V1Handler: V1PetitionHandler, router: router, db: db, AuthMiddleware: KeycloakAuthMiddleware}
}

func (r *PetitionsRoutes) RegisterRoutes() {
	V1PetitionRoute := r.router.Group("/petitions")
	{
		V1PetitionRoute.GET("", middlewares.Pagination(), r.V1Handler.GetAllPetitions)
		V1PetitionRoute.POST("", r.AuthMiddleware(), r.V1Handler.CreatePetition)
		V1PetitionRoute.POST("/:id/like", r.AuthMiddleware(), r.V1Handler.LikePetition)
		V1PetitionRoute.POST("/:id/voice", r.AuthMiddleware(), r.V1Handler.VoicePetition)
		V1PetitionRoute.DELETE("/:id", r.AuthMiddleware(), r.V1Handler.Delete)
	}

}
