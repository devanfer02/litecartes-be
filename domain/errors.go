package domain

import (
    "fmt"
    "strings"
    "errors"
    "net/http"
)

var (
	ErrNotFound         = errors.New("requested item not found")
    ErrConflict         = errors.New("item already exist")
    ErrBadRequest       = errors.New("given request is not valid")
    ErrServerError      = errors.New("internal server error")
    ErrInvalidCredential= errors.New("invalid credential request")
    ErrInvalidToken     = errors.New("invalid token")
    ErrUserRegistered   = errors.New("user already registered")
    ErrUnauthorized     = errors.New("can't acccess resources, unauthorized")
)

func ValidationFailed(message string) error {
    return errors.New(fmt.Sprintf("Validation doesn't pass. Message: %s\n", message))
}

func GetCode(err error) int {
    if err == nil {
        return http.StatusOK
    }

    fmt.Println(err.Error())
    if strings.Contains(err.Error(), "Validation doesn't pass") {
        return http.StatusBadRequest
    }

    switch err {
        case ErrServerError :
            return http.StatusInternalServerError
        case ErrConflict : 
            return http.StatusConflict
        case ErrBadRequest :
            return http.StatusBadRequest
        case ErrNotFound :
            return http.StatusNotFound 
        case ErrInvalidCredential :
            return http.StatusUnauthorized
        case ErrUserRegistered : 
            return http.StatusConflict 
        default :
            return http.StatusInternalServerError
    }
}