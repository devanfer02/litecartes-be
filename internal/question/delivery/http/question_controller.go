package controller

import (
	"github.com/devanfer02/litecartes/domain"
	res "github.com/devanfer02/litecartes/http/response"
	"github.com/devanfer02/litecartes/internal/utils"
	_mdlwr "github.com/devanfer02/litecartes/middleware"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
    queUcase domain.QuestionUsecase
}

func InitQuestionController(qUcase domain.QuestionUsecase, r *gin.Engine) {
    qCtr := &QuestionController{
        queUcase: qUcase,
    }

    qR := r.Group("/questions").Use(_mdlwr.CORS())
    qR.GET("", qCtr.Fetch)
    qR.GET("/:uid", qCtr.FetchByUID)
    qR.POST("", qCtr.CreateQuestion)
    qR.PUT("/:uid", qCtr.UpdateQuestion)
    qR.DELETE("/:uid", qCtr.DeleteQuestion)
}

func(c *QuestionController) Fetch(ctx *gin.Context) {
    pageReq, err := domain.GetPageRequest(ctx)

    if utils.ErrNotNil(ctx, err, domain.GetCode(err)) {
        return 
    }

    questions, resp, err := c.queUcase.Fetch(ctx.Request.Context(), pageReq)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully fetch questions", questions, resp)
}

func(c *QuestionController) FetchByUID(ctx *gin.Context) {
    uidP := ctx.Param("uid")

    question, err := c.queUcase.FetchByUID(ctx.Request.Context(), uidP)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully fetch question", question, nil)
}

func(c *QuestionController) CreateQuestion(ctx *gin.Context) {
    var question domain.Question 

    if utils.BindingFailed(ctx, &question) {
        return 
    }

    err := c.queUcase.Insert(ctx.Request.Context(), &question)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully insert question", question, nil)
}

func(c *QuestionController) UpdateQuestion(ctx *gin.Context) {
    var question domain.Question
    uidP := ctx.Param("uid")

    if utils.BindingFailed(ctx, &question) {
        return 
    }

    question.UID = uidP 

    
    err := c.queUcase.Update(ctx.Request.Context(), &question)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully update question", question, nil)
}

func(c *QuestionController) DeleteQuestion(ctx *gin.Context) {
    uidP := ctx.Param("uid")

    err := c.queUcase.Delete(ctx.Request.Context(), uidP)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully delete question", nil, nil)
}