package models

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Board struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

var DB *sql.DB

func GetBoard(count int) ([]Board, error) {

	rows, err := DB.Query("SELECT id, status, name from boards LIMIT " + strconv.Itoa(count))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	boards := make([]Board, 0)

	for rows.Next() {
		singleBoard := Board{}
		err = rows.Scan(&singleBoard.ID, &singleBoard.Status, &singleBoard.Name)

		if err != nil {
			return nil, err
		}

		boards = append(boards, singleBoard)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return boards, err
}

func GetBoardByID(id string) (Board, error) {
	stmt, err := DB.Prepare("SELECT id, status, name from boards WHERE id = ?")
	if err != nil {
		return Board{}, err
	}

	singleBoard := Board{}

	sqlErr := stmt.QueryRow(id).Scan(&singleBoard.ID, &singleBoard.Status, &singleBoard.Name)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Board{}, nil
		}
		return Board{}, sqlErr
	}
	return singleBoard, nil
}
