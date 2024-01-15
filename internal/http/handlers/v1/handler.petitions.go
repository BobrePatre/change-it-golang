package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	V1DomainErrors "change-it/internal/business/errors/v1"
	"change-it/internal/constants"
	V1Requests "change-it/internal/http/datatransfers/requests/v1"
	V1Responses "change-it/internal/http/datatransfers/responses/v1"
	"change-it/pkg/validators"
	"errors"
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
	petitionDomain.OwnerID = ctx.Value(constants.UserDetails).(V1Domains.UserDetails).UserId
	err := h.usecase.Create(ctx, petitionDomain)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	NewSuccessResponse(ctx, http.StatusOK)
}

func (h *PetitionsHandler) GetAllPetitions(ctx *gin.Context) {

	pageRequest := ctx.Value(constants.PageInfo).(V1Requests.PageRequest)

	petitionDomains, total, err := h.usecase.GetAll(ctx, pageRequest.PageNumber, pageRequest.PageSize)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	outResponses := V1Responses.ArrayFromV1Domains(petitionDomains)
	response := V1Responses.PageResponse{
		CurrentPage: pageRequest.PageNumber,
		PageSize:    pageRequest.PageSize,
		TotalPages:  total,
		Data:        outResponses,
	}
	NewSuccessResponseWithData(ctx, http.StatusOK, response)
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

	userDetails := ctx.Value(constants.UserDetails).(V1Domains.UserDetails)
	err := h.usecase.Delete(ctx, request.ID, userDetails.UserId, userDetails.Roles)
	if err != nil {
		var nferr *V1DomainErrors.NotFoundError
		var ferr *V1DomainErrors.ForbiddenError
		switch {
		case errors.As(err, &nferr):
			NewErrorResponse(ctx, nferr.StatusCode, nferr.Message)
			return
		case errors.As(err, &ferr):
			NewErrorResponse(ctx, ferr.StatusCode, ferr.Message)
			return
		default:
			NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

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

	userId := ctx.Value(constants.UserDetails).(V1Domains.UserDetails).UserId
	err := h.usecase.Like(ctx, request.ID, userId)
	if err != nil {
		var nferr *V1DomainErrors.NotFoundError
		var aerr *V1DomainErrors.AlreadyLikedError
		switch {
		case errors.As(err, &nferr):
			NewErrorResponse(ctx, nferr.StatusCode, nferr.Message)
			return
		case errors.As(err, &aerr):
			NewErrorResponse(ctx, aerr.StatusCode, aerr.Message)
			return
		default:
			NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

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

	userId := ctx.Value(constants.UserDetails).(V1Domains.UserDetails).UserId
	err := h.usecase.Voice(ctx, request.ID, userId)
	if err != nil {
		var nferr *V1DomainErrors.NotFoundError
		var aerr *V1DomainErrors.AlreadyVoicedError
		switch {
		case errors.As(err, &nferr):
			NewErrorResponse(ctx, nferr.StatusCode, nferr.Message)
			return
		case errors.As(err, &aerr):
			NewErrorResponse(ctx, aerr.StatusCode, aerr.Message)
			return
		default:
			NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	NewSuccessResponse(ctx, http.StatusOK)
}
