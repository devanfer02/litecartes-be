package middleware

import (
	"firebase.google.com/go/v4/auth"
	"github.com/devanfer02/litecartes/domain"
)

type Middleware struct {
    userUcase    domain.UserUsecase
    fireAuth     *auth.Client
}

func NewMiddleware(uR domain.UserUsecase, fA *auth.Client) *Middleware {
    return &Middleware{
        userUcase: uR,
        fireAuth: fA,
    }
}