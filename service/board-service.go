package service

import (
	"example.com/kanban/database"
	"example.com/kanban/entity"
)

type BoardService interface {
	Show() ([]entity.Board, error)
	// ShowByID(id string) ([]entity.Board, error)
	Create(entity.Board) (entity.Board, error)
}

type boardService struct {
	boards []entity.Board
}

func NewBoardConstructor() BoardService {
	return &boardService{}
}

func (bs *boardService) Show() ([]entity.Board, error) {
	// Initialize the boards slice if it's nil
	if bs.boards == nil {
		bs.boards = []entity.Board{}
	}

	// Execute the query
	rows, err := database.DB.Query("SELECT id, status, name FROM boards LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var board entity.Board
		err = rows.Scan(&board.ID, &board.Status, &board.Name)
		if err != nil {
			return nil, err
		}
		bs.boards = append(bs.boards, board)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bs.boards, nil
}

// func (bs *boardService) ShowByID(id string) (entity.Board, error) {
// 	stmt, err := database.DB.Prepare("SELECT id, status, name FROM boards WHERE id = ?")
// 	if err != nil {
// 		return nil, err
// 	}

// 	sqlE
// }

func (bs *boardService) Create(board entity.Board) (entity.Board, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return entity.Board{}, err
	}

	stmt, err := tx.Prepare("INSERT INTO boards (status, name) VALUES (?, ?)")
	if err != nil {
		return entity.Board{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(board.Status, board.Name)
	if err != nil {
		return entity.Board{}, err
	}

	tx.Commit()
	return board, nil
}
