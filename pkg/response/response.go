package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success   bool      `json:"success"`
	Message   string    `json:"message"`
	Payload   any       `json:"payload,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Error     *APIError `json:"error,omitempty"`
	Meta      *Meta     `json:"meta,omitempty"`
}

type Meta struct {
	Page        int `json:"page,omitempty"`
	RowsPerPage int `json:"rowsPerPage,omitempty"`
	Total       int `json:"total,omitempty"`
	TotalPages  int `json:"totalPages,omitempty"`
}

type APIError struct {
	Code             int               `json:"code"`
	Message          string            `json:"message,omitempty"`
	Details          []string          `json:"details,omitempty"`
	ValidationErrors map[string]string `json:"validation_errors,omitempty"`
}

func Success(ctx *gin.Context, message string, opts ...Option) {
	res := APIResponse{
		Success:   true,
		Message:   message,
		Timestamp: time.Now(),
	}

	for _, opt := range opts {
		opt(&res)
	}

	ctx.JSON(http.StatusOK, res)
}

func Error(ctx *gin.Context, code int, message string, opts ...Option) {
	res := APIResponse{
		Success:   false,
		Message:   message,
		Timestamp: time.Now(),
		Error: &APIError{
			Code:    code,
			Message: message,
		},
	}

	for _, opt := range opts {
		opt(&res)
	}

	ctx.JSON(code, res)
}
