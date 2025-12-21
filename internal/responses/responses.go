package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse is the standard API response wrapper
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Code    int         `json:"code"`
}

// PaginatedResponse handles paginated data
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// Success returns a success response
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Code:    http.StatusOK,
	})
}

// SuccessCreated returns a 201 created response
func SuccessCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Code:    http.StatusCreated,
	})
}

// SuccessPaginated returns a paginated success response
func SuccessPaginated(c *gin.Context, items interface{}, total int64, page, pageSize int) {
	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	response := APIResponse{
		Success: true,
		Message: "Data retrieved successfully",
		Data: PaginatedResponse{
			Items:      items,
			Total:      total,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
		},
		Code: http.StatusOK,
	}
	c.JSON(http.StatusOK, response)
}

// Error returns an error response
func Error(c *gin.Context, statusCode int, message string, errors interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Message: message,
		Errors:  errors,
		Code:    statusCode,
	})
}

// BadRequest returns 400 Bad Request
func BadRequest(c *gin.Context, message string, errors interface{}) {
	Error(c, http.StatusBadRequest, message, errors)
}

// Unauthorized returns 401 Unauthorized
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message, nil)
}

// Forbidden returns 403 Forbidden
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message, nil)
}

// NotFound returns 404 Not Found
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, nil)
}

// Conflict returns 409 Conflict
func Conflict(c *gin.Context, message string) {
	Error(c, http.StatusConflict, message, nil)
}

// InternalServerError returns 500 Internal Server Error
func InternalServerError(c *gin.Context, message string, err interface{}) {
	Error(c, http.StatusInternalServerError, message, err)
}

// ValidationFailed returns validation error response
func ValidationFailed(c *gin.Context, errors interface{}) {
	Error(c, http.StatusBadRequest, "validation failed", errors)
}

// NoContent returns 204 No Content
func NoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

// Accepted returns 202 Accepted
func Accepted(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusAccepted, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Code:    http.StatusAccepted,
	})
}
