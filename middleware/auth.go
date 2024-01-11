package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Auth() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        headerToken := ctx.GetHeader("Authorization")

        headerToken = strings.Replace(headerToken, "Bearer: ", "", 1)

        token, err := m.fireAuth.VerifyIDToken(ctx, headerToken)
        if err != nil {
            ctx.AbortWithStatus(http.StatusUnauthorized)
        }

        ctx.Set("__userAuthorized", token.UID)
        ctx.Next()
    }
}