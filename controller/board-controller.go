package controller

import (
	"log"
	"net/http"

	"example.com/kanban/entity"
	"example.com/kanban/service"
	"github.com/gin-gonic/gin"
)

type BoardController interface {
	Show() []entity.Board
	// Needs JSON object and should return slice but the previous tutorial return only one
	ShowByID(ctx *gin.Context) entity.Board
	Create(ctx *gin.Context) entity.Board
	Update(ctx *gin.Context) entity.Board
	Delete(ctx *gin.Context) entity.Board
}

type boardController struct {
	service service.BoardService
}

func NewBoardConstructor(service service.BoardService) BoardController {
	return &boardController{
		service: service,
	}
}

func (bc *boardController) Show() []entity.Board {
	boards, err := bc.service.Show()
	if err != nil {
		log.Println(err)
	}
	return boards
}

func (bc *boardController) ShowByID(ctx *gin.Context) entity.Board {
	id := ctx.Param("id")
	boards, err := bc.service.ShowByID(id)
	if err != nil {
		log.Println(err)
	}
	return boards
}

func (bc *boardController) Create(ctx *gin.Context) entity.Board {
	var board entity.Board
	ctx.BindJSON(&board)
	_, err := bc.service.Create(board)
	if err != nil {
		log.Println(err)
	}
	return board
}

func (bc *boardController) Update(ctx *gin.Context) entity.Board {
	var json entity.Board

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return entity.Board{}
	}

	id := ctx.Param("id")

	success, err := bc.service.Update(json, id)
	if success {
		ctx.JSON(http.StatusOK, gin.H{"message": "Board updated successfully"})
		return json
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return entity.Board{}
	}
}

func (bc *boardController) Delete(ctx *gin.Context) entity.Board {
	id := ctx.Param("id")
	boards, err := bc.service.Delete(id)
	if err != nil {
		log.Println(err)
	}
	return boards
}
