package main

import (
	"embed"
	"fmt"
	"net/http"

	"example.com/kanban/controller"
	"example.com/kanban/database"
	migrator "example.com/kanban/database/migration"
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

const migrationsDir = "database/migration"

//go:embed database/migration/*.sql
var MigrationFS embed.FS

func main() {
	database.DBConnection()

	migration := migrator.MustGetNewMigrator(MigrationFS, migrationsDir)
	err := migration.ApplyMigrations(database.DB)
	if err != nil {
		panic(err)
	}

	fmt.Println("Migrations applied!!")

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
	router.DELETE("/boards/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, BoardController.Delete(ctx))
	})

	router.Run("localhost:9000")
}
