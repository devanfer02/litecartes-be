package middleware

import (
	"net/http"

    "github.com/devanfer02/litecartes/bootstrap/env"

	"github.com/gin-gonic/gin"
)

func Interceptor() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        headerKey := ctx.GetHeader("X-API-Key")

        if headerKey != env.ProcEnv.ApiKey {
            ctx.AbortWithStatus(http.StatusForbidden)
            return 
        }

        ctx.Next()
    }
}