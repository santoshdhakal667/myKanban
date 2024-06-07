package models

import (
	_ "github.com/mattn/go-sqlite3"
)

func PostBoard(newBoard Board) (bool, error) {
	//! Add ID check (if it already exists)
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO boards (status, name) VALUES (?,?)")
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newBoard.Status, newBoard.Name)
	if err != nil {
		return false, err
	}

	tx.Commit()
	return true, nil
}
