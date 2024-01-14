package middleware

import (
	"net/http"
	"strings"
    "fmt"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Auth() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        headerToken := ctx.GetHeader("Authorization")

        headerToken = strings.Replace(headerToken, "Bearer ", "", 1)

        token, err := m.fireAuth.VerifyIDToken(ctx, headerToken)
        if err != nil {
            fmt.Println(headerToken)
			fmt.Println(err)
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return 
        }

        _, err = m.userUcase.FetchByUID(ctx.Request.Context(), token.UID)

        if err != nil {
            fmt.Println(err)
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return 
        }

        ctx.Set("__userAuthorized", token.UID)
        ctx.Next()
    }
}