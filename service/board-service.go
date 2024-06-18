package service

import (
	"database/sql"

	"example.com/kanban/database"
	"example.com/kanban/entity"
)

type BoardService interface {
	Show() ([]entity.Board, error)
	ShowByID(id string) (entity.Board, error)
	Create(entity.Board) (entity.Board, error)
}

type boardService struct {
	boards       []entity.Board
	showStmt     *sql.Stmt
	showByIDStmt *sql.Stmt
	createStmt   *sql.Stmt
}

func NewBoardConstructor() (BoardService, error) {
	showStmt, err := database.DB.Prepare("SELECT id, status, name FROM boards LIMIT 100")
	if err != nil {
		return nil, err
	}
	showByIDStmt, err := database.DB.Prepare("SELECT id, status, name FROM boards WHERE id = ?")
	if err != nil {
		return nil, err
	}
	createStmt, err := database.DB.Prepare("INSERT INTO boards (status, name) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	return &boardService{
		boards:       []entity.Board{},
		showStmt:     showStmt,
		showByIDStmt: showByIDStmt,
		createStmt:   createStmt,
	}, nil
}

func (bs *boardService) Show() ([]entity.Board, error) {
	rows, err := bs.showStmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boards []entity.Board
	for rows.Next() {
		var board entity.Board
		err = rows.Scan(&board.ID, &board.Status, &board.Name)
		if err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return boards, nil
}

func (bs *boardService) ShowByID(id string) (entity.Board, error) {
	var board entity.Board
	err := bs.showByIDStmt.QueryRow(id).Scan(&board.ID, &board.Status, &board.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Board{}, nil
		}
		return entity.Board{}, err
	}
	return board, nil
}

func (bs *boardService) Create(board entity.Board) (entity.Board, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return entity.Board{}, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw after rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	_, err = bs.createStmt.Exec(board.Status, board.Name)
	if err != nil {
		return entity.Board{}, err
	}

	return board, nil
}
