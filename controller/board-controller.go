package controller

import (
	"log"

	"example.com/kanban/entity"
	"example.com/kanban/service"
	"github.com/gin-gonic/gin"
)

type BoardController interface {
	Show() []entity.Board
	// Needs JSON object and should return slice but the previous tutorial return only one
	Create(ctx *gin.Context) entity.Board
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

func (bc *boardController) Create(ctx *gin.Context) entity.Board {
	var board entity.Board
	ctx.BindJSON(&board)
	_, err := bc.service.Create(board)
	if err != nil {
		log.Println(err)
	}
	return board
}
