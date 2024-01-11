package utils

import (
    "net/http"

    res "github.com/devanfer02/litecartes/http/response"

    "github.com/gin-gonic/gin"
)

func BindingFailed(ctx *gin.Context, data any) bool {
    if err := ctx.ShouldBindJSON(data); err != nil {
        ctx.JSON(
            http.StatusBadRequest, 
            res.ResponseError{Status: http.StatusBadRequest, Message: err.Error()},
        )
        return true  
    }

    return false 
}

func ErrNotNil(ctx *gin.Context, err error, code int) bool {
    if err != nil {
        ctx.JSON(
            code, 
            res.ResponseError{Status: code, Message: err.Error()},
        )
        return true 
    }

    return false
}