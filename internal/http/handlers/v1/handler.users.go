package v1

import (
	V1Domains "change-it/internal/business/domains/v1"
	"change-it/internal/constants"
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
	outDomains, err := h.usecase.GetLikedPetitions(ctx, userDetails.UserId)
	if err != nil {
		NewAbortResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := V1Resopnses.ArrayFromV1Domains(outDomains)
	NewSuccessResponseWithData(ctx, http.StatusOK, response)

}

func (h *UsersHandler) GetVoicedPetitions(ctx *gin.Context) {

	userDetails := ctx.Value(constants.UserDetails).(V1Domains.UserDetails)
	outDomains, err := h.usecase.GetVoicedPetitions(ctx, userDetails.UserId)
	if err != nil {
		NewAbortResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := V1Resopnses.ArrayFromV1Domains(outDomains)
	NewSuccessResponseWithData(ctx, http.StatusOK, response)
}

func (h *UsersHandler) GetOwnedPetitions(ctx *gin.Context) {

	userDetails := ctx.Value(constants.UserDetails).(V1Domains.UserDetails)
	outDomains, err := h.usecase.GetOwnedPetitions(ctx, userDetails.UserId)
	if err != nil {
		NewAbortResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := V1Resopnses.ArrayFromV1Domains(outDomains)
	NewSuccessResponseWithData(ctx, http.StatusOK, response)

}
