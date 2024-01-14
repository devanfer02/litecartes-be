package middleware 

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func CORS() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowAllOrigins:  true,
        AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions, http.MethodPut},
        AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization", "X-API-Key", "X-Cursor", "Token-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    })
}
