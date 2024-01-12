package controller

import (
	_mdlwr "github.com/devanfer02/litecartes/middleware"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
    
}

func InitTaskController(r *gin.Engine) {
    tCtr := &TaskController{}

    tR := r.Group("/tasks").Use(_mdlwr.CORS())
    tR.GET("", tCtr.Fetch)
    tR.GET("/:uid", tCtr.FetchWithQuestions)
    tR.POST("", tCtr.CreateTask)
    tR.PUT("/:uid", tCtr.UpdateTask)
    tR.DELETE("/:uid", tCtr.DeleteTask)
}

func(c *TaskController) FetchTask(ctx *gin.Context) {

}

func(c *TaskController) FetchWithQuestions(ctx *gin.Context) {
    
}

func(c *TaskController) CreateTask(ctx *gin.Context) {

}

func(c *TaskController) UpdateTask(ctx *gin.Context) {

}

func(c *TaskController) DeleteTask(ctx *gin.Context) {

}