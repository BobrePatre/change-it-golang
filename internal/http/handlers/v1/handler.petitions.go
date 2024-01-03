package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1Requests "change-it/internal/http/datatransfers/requests/v1"
	"change-it/pkg/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PetitionsHandler struct {
	usecase V1Domains.PetitionUseсase
}

func NewPetitionsHandler(usecase V1Domains.PetitionUseсase) PetitionsHandler {
	return PetitionsHandler{
		usecase: usecase,
	}
}

func (h *PetitionsHandler) CreatePetition(ctx *gin.Context) {

	var request V1Requests.PetitionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	petitionDomain := request.ToV1Domain()
	err := h.usecase.Create(ctx, petitionDomain)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, http.StatusOK)
}
