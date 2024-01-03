package v1

import (
	V1Usecase "change-it/internal/business/usecases/v1"
	V1PostgresRepository "change-it/internal/datasources/repositories/postgres/v1"
	V1Handler "change-it/internal/http/handlers/v1"
	V1Handlers "change-it/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type PetitionsRoutes struct {
	V1Handler V1Handlers.PetitionsHandler
	router    *gin.RouterGroup
	db        *sqlx.DB
}

func NewPetitionRoute(router *gin.RouterGroup, db *sqlx.DB) *PetitionsRoutes {
	V1PetitionRepository := V1PostgresRepository.NewPetitionRepository(db)
	V1PetitionUsecase := V1Usecase.NewPetitionUsecase(V1PetitionRepository)
	V1PetitionHandler := V1Handler.NewPetitionsHandler(V1PetitionUsecase)

	return &PetitionsRoutes{V1Handler: V1PetitionHandler, router: router, db: db}
}

func (r *PetitionsRoutes) RegisterRoutes() {
	V1Route := r.router.Group("/v1")
	{

		V1PetitionRoute := V1Route.Group("/petitions")
		V1PetitionRoute.POST("/", r.V1Handler.CreatePetition)
		V1PetitionRoute.GET("/", r.V1Handler.GetAllPetitions)
		V1PetitionRoute.DELETE("/:id", r.V1Handler.Delete)
		V1PetitionRoute.POST("/:id/like", r.V1Handler.LikePetition)
		V1PetitionRoute.POST("/:id/voice", r.V1Handler.VoicePetition)

	}
}
