package model

import (
	"strconv"

	"example.com/kanban/database"
	_ "github.com/go-sql-driver/mysql"
)

type Board struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

func GetBoard(count int) ([]Board, error) {
	rows, err := database.DB.Query("SELECT id, status, name FROM boards LIMIT " + strconv.Itoa(count))
	// rows, err := DB.Query("SELECT id, status, name FROM boards LIMIT " + strconv.Itoa(count))
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

	return boards, err
}

// func GetBoardByID(id string) (Board, error) {
// 	stmt, err := DB.Prepare("SELECT id, status, name from boards WHERE id = ?")
// 	if err != nil {
// 		return Board{}, err
// 	}

// 	singleBoard := Board{}

// 	sqlErr := stmt.QueryRow(id).Scan(&singleBoard.ID, &singleBoard.Status, &singleBoard.Name)

// 	if sqlErr != nil {
// 		if sqlErr == sql.ErrNoRows {
// 			return Board{}, nil
// 		}
// 		return Board{}, sqlErr
// 	}
// 	return singleBoard, nil
// }

// func DeleteBoard(boardID int) (bool, error) {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return false, err
// 	}

// 	stmt, err := tx.Prepare("DELETE FROM boards WHERE id = ?")
// 	if err != nil {
// 		return false, err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(boardID)
// 	if err != nil {
// 		return false, err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// func PostBoard(newBoard Board) (bool, error) {
// 	//?? Add ID check (if it already exists)
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return false, err
// 	}

// 	stmt, err := tx.Prepare("INSERT INTO boards (status, name) VALUES (?,?)")
// 	if err != nil {
// 		return false, err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(newBoard.Status, newBoard.Name)
// 	if err != nil {
// 		return false, err
// 	}

// 	tx.Commit()
// 	return true, nil
// }

// func PutBoard(board Board, id int) (bool, error) {
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		return false, err
// 	}

// 	stmt, err := tx.Prepare("UPDATE boards SET status =?, name =? WHERE id =?")
// 	if err != nil {
// 		return false, err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(board.Status, board.Name, board.ID)
// 	if err != nil {
// 		return false, err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
