package models

import (
	_ "github.com/mattn/go-sqlite3"
)

func PutBoard(board Board, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE boards SET status =?, name =? WHERE id =?")
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(board.Status, board.Name, board.ID)
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}
