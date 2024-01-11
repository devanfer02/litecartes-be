package response

import (
    "github.com/devanfer02/litecartes/domain"

    "github.com/gin-gonic/gin"
)

type ResponseError struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
}

type Response struct {
	Status     int                          `json:"status"`
	Message    string                       `json:"message"`
	Data       interface{}                  `json:"data,omitempty"`
	Pagination *domain.PaginationResponse   `json:"pagination,omitempty"`
}

func SendResponse(
    ctx *gin.Context, 
    code int, 
    message string, 
    data interface{}, 
    pagination *domain.PaginationResponse,
    ) {
    
    ctx.JSON(code, Response{Status: code, Message: message, Data: data, Pagination: pagination})
}