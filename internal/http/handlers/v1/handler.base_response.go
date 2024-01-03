package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewSuccessResponse(c *gin.Context, statusCode int) {
	c.Status(statusCode)
}

func NewSuccessResponseWithData(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

func NewErrorResponse(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{"message": err})

}

func NewAbortResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": message})
}
