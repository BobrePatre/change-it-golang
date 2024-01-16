package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	"change-it/internal/constants"
	V1Requests "change-it/internal/http/datatransfers/requests/v1"
	V1Resopnses "change-it/internal/http/datatransfers/responses/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersHandler struct {
	usecase V1Domains.UserUsecase
}

func NewUsersHandler(usecase V1Domains.UserUsecase) UsersHandler {
	return UsersHandler{
		usecase: usecase,
	}
}

func (h *UsersHandler) GetLikedPetitions(ctx *gin.Context) {

	userDetails := ctx.Value(constants.UserDetails).(V1Domains.UserDetails)
	pageRequest := ctx.Value(constants.PageInfo).(V1Requests.PageRequest)
	outDomains, total, err := h.usecase.GetLikedPetitions(ctx, userDetails.UserId, pageRequest.PageNumber, pageRequest.PageSize)
	if err != nil {
		NewAbortResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	outResponses := V1Resopnses.ArrayFromV1Domains(outDomains)
	response := V1Resopnses.PageResponse{
		Data:        outResponses,
		PageSize:    pageRequest.PageSize,
		CurrentPage: pageRequest.PageNumber,
		TotalPages:  total,
	}
	NewSuccessResponseWithData(ctx, http.StatusOK, response)

}

func (h *UsersHandler) GetVoicedPetitions(ctx *gin.Context) {

	userDetails := ctx.Value(constants.UserDetails).(V1Domains.UserDetails)
	pageRequest := ctx.Value(constants.PageInfo).(V1Requests.PageRequest)
	outDomains, total, err := h.usecase.GetVoicedPetitions(ctx, userDetails.UserId, pageRequest.PageNumber, pageRequest.PageSize)
	if err != nil {
		NewAbortResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	outResponses := V1Resopnses.ArrayFromV1Domains(outDomains)
	response := V1Resopnses.PageResponse{
		Data:        outResponses,
		PageSize:    pageRequest.PageSize,
		CurrentPage: pageRequest.PageNumber,
		TotalPages:  total,
	}
	NewSuccessResponseWithData(ctx, http.StatusOK, response)
}

func (h *UsersHandler) GetOwnedPetitions(ctx *gin.Context) {

	userDetails := ctx.Value(constants.UserDetails).(V1Domains.UserDetails)
	pageRequest := ctx.Value(constants.PageInfo).(V1Requests.PageRequest)
	outDomains, total, err := h.usecase.GetOwnedPetitions(ctx, userDetails.UserId, pageRequest.PageNumber, pageRequest.PageSize)
	if err != nil {
		NewAbortResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	outResponses := V1Resopnses.ArrayFromV1Domains(outDomains)
	response := V1Resopnses.PageResponse{
		Data:        outResponses,
		PageSize:    pageRequest.PageSize,
		CurrentPage: pageRequest.PageNumber,
		TotalPages:  total,
	}
	NewSuccessResponseWithData(ctx, http.StatusOK, response)

}
