package errors

import (
	"net/http"

	resp "github.com/cuonglv-smartosc/golang-boiler-template/pkg/response"
	"github.com/gin-gonic/gin"
)

func respondError(c *gin.Context, status int, msg string, code string, details []ErrorDetail) {
	if details != nil {
		resp.FailWithCodeAndDetail(c, status, msg, code, convertErrorDetails(details))
		return
	}
	if code != "" {
		resp.FailWithCode(c, status, msg, code)
		return
	}
	resp.Fail(c, status, msg)
}

func HandleValidationError(c *gin.Context, err error) bool {
	if err != nil {
		errorMsg, details := FormatValidationErrors(err)
		respondError(c, http.StatusBadRequest, errorMsg, "VALIDATION_ERROR", details)
		return true
	}
	return false
}

func HandleValidationErrorWithCode(c *gin.Context, err error, code string) bool {
	if err != nil {
		errorMsg, details := FormatValidationErrors(err)
		respondError(c, http.StatusBadRequest, errorMsg, code, details)
		return true
	}
	return false
}

func HandleCustomError(c *gin.Context, err error, code string, message string) {
	var status int
	var errorMsg string
	if err != nil {
		serviceStatus, serviceMsg, mappedCode, details := MapServiceError(err)
		status = serviceStatus
		if message != "" {
			errorMsg = message
		} else {
			errorMsg = serviceMsg
		}
		if code == "" {
			code = mappedCode
		}
		respondError(c, status, errorMsg, code, details)
		return
	}
	customErr := NewCustomError(422, code, message, nil)
	businessStatus, businessMsg, businessCode, details := MapServiceError(customErr)
	respondError(c, businessStatus, businessMsg, businessCode, details)
}

func HandleServerError(c *gin.Context, message string) {
	resp.FailWithCode(c, 500, message, "500")
}

func HandleForbiddenError(c *gin.Context, message string) {
	resp.FailWithCode(c, 403, message, "403")
}

func HandleNotFoundError(c *gin.Context, message string) {
	resp.FailWithCode(c, 404, message, "404")
}

func HandleRateLimitError(c *gin.Context, message string) {
	resp.FailWithCode(c, 429, message, "429")
}

func convertErrorDetails(errDetails []ErrorDetail) []resp.ErrorDetail {
	var respDetails []resp.ErrorDetail
	for _, detail := range errDetails {
		respDetails = append(respDetails, resp.ErrorDetail{
			Code:    detail.Code,
			Message: detail.Message,
			Field:   detail.Field,
		})
	}
	return respDetails
}
