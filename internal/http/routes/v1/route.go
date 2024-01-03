package v1

import (
	V1Handler "change-it/internal/http/handlers/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootHandler(ctx *gin.Context) {
	V1Handler.NewSuccessResponseWithData(ctx, http.StatusOK, gin.H{"message": "Welcome to change-it API"})
}
