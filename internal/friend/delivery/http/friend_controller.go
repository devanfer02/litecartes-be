package controller

import (
	"github.com/devanfer02/litecartes/domain"
    "github.com/devanfer02/litecartes/internal/utils"
    res "github.com/devanfer02/litecartes/http/response"
    _mdlwr "github.com/devanfer02/litecartes/middleware"

	"github.com/gin-gonic/gin"
)

type FriendController struct {
	friendUcase domain.FriendUsecase
}

func InitFriendController(firendUcase domain.FriendUsecase, mdlwr *_mdlwr.Middleware,  r *gin.Engine) {
    fCtr := &FriendController{friendUcase: firendUcase}

    fR := r.Group("/friends").Use(_mdlwr.CORS())
    fR.GET("/followers", mdlwr.Auth(), fCtr.FetchFollowers)
    fR.GET("/followings", mdlwr.Auth(), fCtr.FetchFollowings)
    fR.POST("/followings/:followedid", mdlwr.Auth(), fCtr.FollowUser)
    fR.POST("/:followedid", mdlwr.Auth(), fCtr.RemoveFriend)
}

func(c *FriendController) FetchFollowers(ctx *gin.Context) {
    uid := ctx.GetString("__userAuthorized")

    users, err := c.friendUcase.FetchFollowers(ctx.Request.Context(), uid)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully fetch followers", users, nil)
}

func(c *FriendController) FetchFollowings(ctx *gin.Context) {
    uid := ctx.GetString("__userAuthorized")

    users, err := c.friendUcase.FetchFollowings(ctx.Request.Context(), uid)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully fetch followers", users, nil)
}

func(c *FriendController) FollowUser(ctx *gin.Context) {
    uid := ctx.GetString("__userAuthorized")
    followedID := ctx.Param("followedid")

    err := c.friendUcase.InsertNewFollower(ctx.Request.Context(), followedID, uid)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully follow user", nil, nil)
}

func(c *FriendController) RemoveFriend(ctx *gin.Context) {
    uid := ctx.GetString("__userAuthorized")
    followedID := ctx.Param("followedid")

    err := c.friendUcase.DeleteFriend(ctx.Request.Context(), followedID, uid)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully remove friend", nil, nil)
}

