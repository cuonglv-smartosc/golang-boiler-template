package errors

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

// Common domain errors
type AuthError struct {
	Code    string
	Message string
}

func (e *AuthError) Error() string { return e.Message }

func NewAuthError(code, message string) *AuthError {
	return &AuthError{Code: code, Message: message}
}

type NotFoundError struct {
	Resource string
	Value    string
}

func (e *NotFoundError) Error() string { return fmt.Sprintf("%s not found: %s", e.Resource, e.Value) }

func NewNotFoundError(resource, value string) *NotFoundError {
	return &NotFoundError{Resource: resource, Value: value}
}

type UnauthorizedError struct {
	Action string
	Reason string
}

func (e *UnauthorizedError) Error() string {
	if e.Reason != "" {
		return fmt.Sprintf("unauthorized: %s - %s", e.Action, e.Reason)
	}
	return fmt.Sprintf("unauthorized: %s", e.Action)
}

func NewUnauthorizedError(action, reason string) *UnauthorizedError {
	return &UnauthorizedError{Action: action, Reason: reason}
}

// Business specific errors
type BusinessLogicError struct {
	Code    string
	Message string
}

func (e *BusinessLogicError) Error() string { return e.Message }

func NewBusinessLogicError(code, message string) *BusinessLogicError {
	return &BusinessLogicError{Code: code, Message: message}
}

// CustomError covers non-typical cases with explicit HTTP status and optional details
type CustomError struct {
	Status  int
	Code    string
	Message string
	Details []ErrorDetail
}

func (e *CustomError) Error() string { return e.Message }

func NewCustomError(status int, code, message string, details []ErrorDetail) *CustomError {
	return &CustomError{Status: status, Code: code, Message: message, Details: details}
}

// FormatValidationErrors maps validator errors to message and details list
func FormatValidationErrors(result error) (string, []ErrorDetail) {
	if validationErrors, ok := result.(validator.ValidationErrors); ok {
		var errorMessages []string
		var details []ErrorDetail

		for _, fieldError := range validationErrors {
			var message string
			switch fieldError.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", fieldError.Field())
			case "email":
				message = fmt.Sprintf("%s must be a valid email address", fieldError.Field())
			case "min":
				message = fmt.Sprintf("%s must be at least %s characters long", fieldError.Field(), fieldError.Param())
			case "max":
				message = fmt.Sprintf("%s must be at most %s characters long", fieldError.Field(), fieldError.Param())
			case "len":
				message = fmt.Sprintf("%s must be exactly %s characters long", fieldError.Field(), fieldError.Param())
			default:
				message = fmt.Sprintf("%s: %s", fieldError.Field(), fieldError.Tag())
			}

			errorMessages = append(errorMessages, message)
			details = append(details, ErrorDetail{
				Field:   fieldError.Field(),
				Message: message,
				Code:    strings.ToUpper(fieldError.Tag()),
			})
		}

		return strings.Join(errorMessages, "; "), details
	}
	return "Invalid request body format", nil
}

// MapServiceError maps domain errors to (status, message, code, details)
func MapServiceError(err error) (int, string, string, []ErrorDetail) {
	if err == nil {
		return 200, "", "", nil
	}
	switch v := err.(type) {
	case *NotFoundError:
		details := []ErrorDetail{{
			Code:    "NOT_FOUND",
			Message: v.Error(),
			Field:   "resource",
		}}
		return 404, "Resource not found", "404", details

	case *AuthError:
		details := []ErrorDetail{{
			Code:    v.Code,
			Message: v.Message,
		}}
		return 401, "Authentication failed", "401", details

	case *UnauthorizedError:
		details := []ErrorDetail{{
			Code:    "UNAUTHORIZED",
			Message: v.Error(),
		}}
		return 401, "Unauthorized", "401", details

	case *BusinessLogicError:
		details := []ErrorDetail{{
			Code:    v.Code,
			Message: v.Message,
		}}
		return 422, "Business logic error", v.Code, details

	case *CustomError:
		return v.Status, v.Message, v.Code, v.Details

	default:
		return 500, "Internal server error", "500", nil
	}
}
