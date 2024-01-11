package domain

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Cursor struct {
    CreatedAt   string     `json:"created_at"`
    PointNext   bool        `json:"poin_next"`
    LimitData   int         `json:"limit_data"`
}

type PaginationRequest struct {
    Limit           int 
    PointNext       bool 
    EncodedCursor   string 
}

type PaginationResponse struct {
    PrevCursor          string `json:"prev_cursor"`
    NextCursor          string `json:"next_cursor"`
    LimitData           int    `json:"limit"`
}

func GetPageRequest(ctx *gin.Context) (*PaginationRequest, error) {
    limitQ := ctx.Query("limit")
    cursorQ := ctx.Query("cursor")
    nextQ := ctx.Query("next")

    var nextB bool

    switch nextQ {
        case "false" :
            nextB = false
        default : 
            nextB = true 
    }

    var limitI int
    var err error

    if limitQ == "" {
        limitI = 10
    } else {
        limitI, err = strconv.Atoi(limitQ)
        if err != nil {
            return nil, ErrBadRequest
        }
    }



    return  &PaginationRequest{
        Limit: limitI,
        PointNext: nextB,
        EncodedCursor: cursorQ,
    }, nil 
}