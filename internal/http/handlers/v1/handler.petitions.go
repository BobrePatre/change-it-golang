package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1DomainErrors "change-it/internal/business/errors/v1"
	V1Requests "change-it/internal/http/datatransfers/requests/v1"
	V1Responses "change-it/internal/http/datatransfers/responses/v1"
	"change-it/pkg/validators"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	var request V1Requests.CreatePetition
	if err := ctx.ShouldBindJSON(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	petitionDomain := request.ToV1Domain()
	err := h.usecase.Save(ctx, petitionDomain)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	NewSuccessResponse(ctx, http.StatusOK)
}

func (h *PetitionsHandler) GetAllPetitions(ctx *gin.Context) {
	petitionDomains, err := h.usecase.GetAll(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponseWithData(ctx, http.StatusOK, V1Responses.ArrayFromV1Domains(petitionDomains))
}

func (h *PetitionsHandler) Delete(ctx *gin.Context) {
	var request V1Requests.Id

	if err := ctx.ShouldBindUri(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.usecase.Delete(ctx, request.ID, uuid.New().String())
	if err != nil {
		var nferr *V1DomainErrors.NotFoundError
		if errors.As(err, &nferr) {
			NewErrorResponse(ctx, 409, err.Error()+string(rune(nferr.StatusCode)))
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessResponse(ctx, http.StatusOK)
}

func (h *PetitionsHandler) LikePetition(ctx *gin.Context) {

	var request V1Requests.Id
	if err := ctx.ShouldBindUri(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.usecase.Like(ctx, request.ID, uuid.New().String())
	if err != nil {
		var nferr *V1DomainErrors.NotFoundError
		if errors.As(err, &nferr) {
			NewErrorResponse(ctx, nferr.StatusCode, err.Error())
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessResponse(ctx, http.StatusOK)
}

func (h *PetitionsHandler) VoicePetition(ctx *gin.Context) {
	var request V1Requests.Id
	if err := ctx.ShouldBindUri(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.usecase.Voice(ctx, request.ID, uuid.New().String())
	if err != nil {
		var nferr *V1DomainErrors.NotFoundError
		if errors.As(err, &nferr) {
			NewErrorResponse(ctx, nferr.StatusCode, err.Error())
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessResponse(ctx, http.StatusOK)
}
