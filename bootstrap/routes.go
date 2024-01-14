package bootstrap

import (
    "net/http"
	"database/sql"
	"time"

    _firebase "github.com/devanfer02/litecartes/bootstrap/firebase"

    _mdlwr "github.com/devanfer02/litecartes/middleware"
    
	_userCtr "github.com/devanfer02/litecartes/internal/user/delivery/http"
	_userRepo "github.com/devanfer02/litecartes/internal/user/repository/mysql"
	_userUcase "github.com/devanfer02/litecartes/internal/user/usecase"

    _questCtr "github.com/devanfer02/litecartes/internal/question/delivery/http"
	_questRepo "github.com/devanfer02/litecartes/internal/question/repository/mysql"
	_questUcase "github.com/devanfer02/litecartes/internal/question/usecase"
    
    _taskCtr "github.com/devanfer02/litecartes/internal/task/delivery/http"
	_taskRepo "github.com/devanfer02/litecartes/internal/task/repository/mysql"
	_taskUcase "github.com/devanfer02/litecartes/internal/task/usecase"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, sqldb *sql.DB) {
    r.GET("/health", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{
            "message": "server running",
        })
    })

    fireAuth := _firebase.GetAuthClient()
    

    userRepo := _userRepo.NewMysqlUserRepository(sqldb)
    userUcase := _userUcase.NewUserUsecase(userRepo, 20 * time.Second, fireAuth)

    questRepo := _questRepo.NewMysqlQuestionRepository(sqldb)
    questUcase := _questUcase.NewQuestionUsecase(questRepo, 20 * time.Second)

    taskRepo := _taskRepo.NewMysqlTaskRepository(sqldb)
    taskUcase := _taskUcase.NewTaskUsecase(taskRepo, questRepo, 20 * time.Second)

    mdlwr := _mdlwr.NewMiddleware(userUcase, fireAuth)

    _userCtr.InitUserController(userUcase, r)
    _questCtr.InitQuestionController(questUcase, r)
    _taskCtr.InitTaskController(taskUcase, mdlwr, r)
    
}   