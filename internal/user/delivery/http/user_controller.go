package controller

import (
	"github.com/devanfer02/litecartes/domain"
	res "github.com/devanfer02/litecartes/http/response"
	"github.com/devanfer02/litecartes/internal/utils"
	_mdlwr "github.com/devanfer02/litecartes/middleware"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    userUcase     domain.UserUsecase
}

func InitUserController(ucase domain.UserUsecase, r *gin.Engine) {
    uCtr := &UserController{
        userUcase: ucase,
    }

    //BELOM KE TESTING SEMUA
    uR := r.Group("/users").Use(_mdlwr.CORS())   
    uR.GET("", uCtr.Fetch)
    uR.GET("/:uid", uCtr.FetchByUID)
    uR.GET("/username/:username", uCtr.FetchUsersByUsername)
    uR.POST("/:uid", uCtr.RegisterUser)
    uR.PUT("/:uid", uCtr.UpdateUser)
    uR.DELETE("/:uid", uCtr.DeleteUser)
}

func(c *UserController) Fetch(ctx *gin.Context) {
    pageReq, err := domain.GetPageRequest(ctx)

    if utils.ErrNotNil(ctx, err, domain.GetCode(err)) {
        return 
    }
    
    users, resp, err := c.userUcase.Fetch(ctx.Request.Context(), pageReq)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code ) {
        return 
    }

    res.SendResponse(ctx, code, "successfully fetch users", users, resp)
}

func(c *UserController) FetchByUID(ctx *gin.Context) {
    idP := ctx.Param("uid")

    user, err := c.userUcase.FetchByUID(ctx.Request.Context(), idP)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully fech users", user, nil)
}

func(c *UserController) FetchUsersByUsername(ctx *gin.Context) {
    username := ctx.Param("username")

    users, err := c.userUcase.FetchUsersByUsername(ctx.Request.Context(), username)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully fetch users by username", users, nil)
}

func(c *UserController) RegisterUser(ctx *gin.Context) {
    uidParam := ctx.Param("uid")

    err := c.userUcase.RegisterUser(ctx.Request.Context(), uidParam)    
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully register user", nil, nil)
}

func(c *UserController) UpdateUser(ctx *gin.Context) {
    var payload domain.UserUpdate
    uidParam := ctx.Param("uid")

    if utils.BindingFailed(ctx, &payload) {
        return 
    }

    payload.UID = uidParam
    err := c.userUcase.UpdateUser(ctx.Request.Context(), &payload)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully update user", nil, nil)
}

func(c *UserController) DeleteUser(ctx *gin.Context) {
    uidParam := ctx.Param("uid")
    
    err := c.userUcase.DeleteUser(ctx.Request.Context(), uidParam)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully delete user", nil, nil)
}