package main

import (
	"fmt"
	"os"

	"github.com/devanfer02/litecartes/bootstrap"
	"github.com/devanfer02/litecartes/bootstrap/database/mysql"
	"github.com/devanfer02/litecartes/bootstrap/env"

	mdlwr "github.com/devanfer02/litecartes/middleware"

	"github.com/gin-gonic/gin"
)

func main() {    
    mysqldb := mysql.NewMysqlConn()
    defer mysqldb.Close()

    if len(os.Args) > 1 && os.Args[1] == "generate" {
        mysql.GenerateSeeders(mysqldb)
    }

    app := gin.Default()
    app.Use(mdlwr.CORS())

    bootstrap.InitRoutes(app, mysqldb)


    app.Run(fmt.Sprintf(":%s", env.ProcEnv.ServerAddress))
}