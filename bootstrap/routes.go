package bootstrap

import (
    "net/http"
	"database/sql"
	"time"

    _firebase "github.com/devanfer02/litecartes/bootstrap/firebase"
    
	_userCtr "github.com/devanfer02/litecartes/internal/user/delivery/http"
	_userRepo "github.com/devanfer02/litecartes/internal/user/repository/mysql"
	_userUcase "github.com/devanfer02/litecartes/internal/user/usecase"

    _questCtr "github.com/devanfer02/litecartes/internal/question/delivery/http"
	_questRepo "github.com/devanfer02/litecartes/internal/question/repository/mysql"
	_questUcase "github.com/devanfer02/litecartes/internal/question/usecase"
    
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
    questUcase := _questUcase.NewQuestionUsecase(questRepo, 12 * time.Second)

    _userCtr.InitUserController(userUcase, r)
    _questCtr.InitQuestionController(questUcase, r)
    
}   