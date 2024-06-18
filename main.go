package main

import (
	"net/http"

	"example.com/kanban/controller"
	"example.com/kanban/database"
	"example.com/kanban/service"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var (
	boardService    service.BoardService       = service.NewBoardConstructor()
	BoardController controller.BoardController = controller.NewBoardConstructor(boardService)
)

func main() {
	database.DBConnection()
	router := gin.Default()

	router.GET("/boards", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, BoardController.Show())
	})
	router.GET("/boards/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, BoardController.ShowByID(ctx))
	})
	router.POST("/boards", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, BoardController.Create(ctx))
	})
	router.PUT("/boards/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, BoardController.Update(ctx))
	})

	router.Run("localhost:9000")
}
