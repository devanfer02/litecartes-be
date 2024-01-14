package controller

import (
	"github.com/devanfer02/litecartes/domain"
    res "github.com/devanfer02/litecartes/http/response"
	"github.com/devanfer02/litecartes/internal/utils"
	_mdlwr "github.com/devanfer02/litecartes/middleware"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
    taskUcase domain.TaskUsecase
}

func InitTaskController(
    taskUcase domain.TaskUsecase, 
    mdlwr *_mdlwr.Middleware,
    r *gin.Engine,
) {
    tCtr := &TaskController{
        taskUcase,
    }

    tR := r.Group("/tasks").Use(_mdlwr.CORS())
    tR.GET("", tCtr.FetchAllTask)
    tR.GET("/users/:uid", tCtr.FetchTasksByUserUID)
    tR.GET("/:uid", tCtr.FetchTaskQuestions)
    tR.POST("", tCtr.CreateTask)
    tR.POST("/completed/:taskuid", mdlwr.Auth(), tCtr.UpdateCompletedTask)
    
    tR.PUT("/:uid", tCtr.UpdateTask)
    tR.DELETE("/:uid", tCtr.DeleteTask)
}

func(c *TaskController) FetchAllTask(ctx *gin.Context) {
    pageReq, err := domain.GetPageRequest(ctx)

    if utils.ErrNotNil(ctx, err, domain.GetCode(err)) {
        return 
    }

    tasks,resp, err := c.taskUcase.FetchAll(ctx.Request.Context(), pageReq)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully fetch all tasks", tasks, resp)
}

func(c *TaskController) FetchTasksByUserUID(ctx *gin.Context) {
    pageReq, err := domain.GetPageRequest(ctx)
    userUID := ctx.Param("uid")

    if utils.ErrNotNil(ctx, err, domain.GetCode(err)) {
        return
    }

    tasks, resp, err := c.taskUcase.FetchTasksByUserUID(ctx.Request.Context(), pageReq, userUID)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return
    }
    
    res.SendResponse(ctx, code, "successfully fetch tasks by user uid", tasks, resp)
}

func(c *TaskController) FetchTaskQuestions(ctx *gin.Context) {
    taskUID := ctx.Param("uid")

    task, err := c.taskUcase.FetchTaskWithQuestions(ctx.Request.Context(), taskUID)    
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully task with questions", task, nil)
}

func(c *TaskController) CreateTask(ctx *gin.Context) {
    var task domain.Task 

    if utils.BindingFailed(ctx, &task) {
        return 
    }

    err := c.taskUcase.InsertTask(ctx, &task)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully create task", task, nil)
}

func(c *TaskController) UpdateCompletedTask(ctx *gin.Context) {
    var completedTask domain.CompletedTask

    completedTask.TaskUID = ctx.Param("taskuid")
    
    completedTask.UserUID = ctx.GetString("__userAuthorized")

    if completedTask.UserUID == "" {
        res.SendResponse(ctx, 401, "request header authorization empty", nil, nil)
        return 
    }

    err := c.taskUcase.InsertCompletedTask(ctx.Request.Context(), &completedTask)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully update user's completed task", nil, nil)
}

func(c *TaskController) UpdateTask(ctx *gin.Context) {
    var task domain.Task 
    task.UID = ctx.Param("uid")

    if utils.BindingFailed(ctx, &task) {
        return 
    }

    err := c.taskUcase.UpdateTask(ctx.Request.Context(), &task)
    code := domain.GetCode(err)

    if utils.ErrNotNil(ctx, err, code) {
        return
    }

    res.SendResponse(ctx, code, "successfully update task", task, nil)
}   

func(c *TaskController) DeleteTask(ctx *gin.Context) {
    taskUID := ctx.Param("uid")

    err := c.taskUcase.DeleteTask(ctx.Request.Context(), taskUID)
    code := domain.GetCode(err)
    
    if utils.ErrNotNil(ctx, err, code) {
        return 
    }

    res.SendResponse(ctx, code, "successfully delete task", nil, nil)
}