package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UnauthorizedResp(ctx *gin.Context, error string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"status":        "error",
		"error_message": error,
	})
}

func BadRequestResp(ctx *gin.Context, error string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"status":        "BadRequest",
		"error_message": error,
	})
}

func ForbiddenResp(ctx *gin.Context, error string) {
	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"status":        "Forbidden",
		"error_message": error,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, error string) {
	ctx.AbortWithStatusJSON(ANZRSP(error), gin.H{
		"status":        "error",
		"error_message": error,
	})
}

func ANZRSP(body string) int {
	if strings.Contains(body, "401") {
		return http.StatusUnauthorized
	} else if strings.Contains(body, "400") {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError

}
