package middlewares

import (
	"change-it/internal/constants"
	V1Requests "change-it/internal/http/datatransfers/requests/v1"
	"change-it/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var pRequest V1Requests.PageRequest
		err := ctx.ShouldBindQuery(&pRequest)
		if err != nil {
			logger.Error("sheet of pagination", nil)

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if pRequest.PageSize == 0 {
			pRequest.PageSize = constants.DefaultSize
		}

		if pRequest.PageNumber == 0 {
			pRequest.PageNumber = constants.DefaultPage
		}

		ctx.Set(constants.PageInfo, pRequest)

		ctx.Next()
	}
}
