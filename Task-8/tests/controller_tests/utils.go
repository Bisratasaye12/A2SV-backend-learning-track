package tests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type AppError struct {
	Err error
	StatusCode int
}

func (appError *AppError) Error() string {
	return appError.Err.Error()
}

type AppResult struct {
	Data interface{}
	Message string
	Err error
	StatusCode int
}

type appHandler func(ctx *gin.Context) *AppResult




func ServeHTTP(handle appHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := handle(ctx)
		if result == nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "INTERNAL SERVER ERROR",
				Data: nil,
			})
		}

		if result.Err == nil {
			ctx.JSON(result.StatusCode, Response{
				Success: true,
				Message: result.Message,
				Data: result.Data,
			})
		} else {
			ctx.JSON(result.StatusCode, Response{
				Success: false,
				Message: result.Err.Error(),
				Data: result.Data,
			})
		}
	}
}