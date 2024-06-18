// Package handlers manages HTTP requests.
package handlers

import (
	"log"
	"net/http"

	model "example.com/kanban/model"
	"github.com/gin-gonic/gin"
)

// GetBoardByID handles retrieving a specific board by its ID.
//
// It expects the following:
// - The board ID as a URL parameter.
//
// The function returns a JSON response with the board data or an error message.
func GetBoardByID(c *gin.Context) {
	// Retrieve the board ID from the URL parameter.
	id := c.Param("id")

	// Retrieve the board from the database using the provided ID.
	board, err := model.GetBoardByID(id)
	if err != nil {
		// If retrieval fails, log the error and terminate the application.
		log.Fatal(err)
	}

	// If the board's name is blank, assume no record is found.
	if board.Name == "" {
		// Respond with a 400 Bad Request status and a no records found message.
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	}

	// If the board is found, respond with a 200 OK status and the board data.
	c.JSON(http.StatusOK, gin.H{"data": board})
}

// // PostBoard handles the creation of a new board.
// //
// // It expects the following:
// // - A JSON body containing the fields for the new Board struct.
// //
// // The function returns a JSON response with the status and a message indicating the result.
// func PostBoard(c *gin.Context) {
// 	// Initialize a variable to hold the JSON body parsed into a Board struct.
// 	var json model.Board

// 	// Parse the JSON request body into the Board struct.
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		// If parsing fails, respond with a 400 Bad Request status and the error message.
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		// Log the error for server-side debugging.
// 		log.Println(err.Error())
// 		return
// 	}

// 	// Log the successfully parsed Board struct for debugging purposes.
// 	log.Println(json)

// 	// Attempt to create the new board in the database using the provided data.
// 	success, err := model.PostBoard(json)
// 	if success {
// 		// If the creation is successful, respond with a 200 OK status and a success message.
// 		c.JSON(http.StatusOK, gin.H{"success": success})
// 	} else {
// 		// If the creation fails, respond with a 400 Bad Request status and the error message.
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 		// Log the error for server-side debugging.
// 		log.Println(err)
// 	}
// }

// /*
// PutBoard handles the update of a board.

// It expects the following:
// - A JSON body containing the updated fields for the Board struct.
// - The board ID as a URL parameter.

// The function returns a JSON response with the status and a message indicating the result.
// */
func PutBoard(c *gin.Context) {
	// Initialize a variable to hold the JSON body parsed into a Board struct.
	var json model.Board

	// Parse the JSON request body into the Board struct.
	if err := c.ShouldBindJSON(&json); err != nil {
		// If parsing fails, respond with a 400 Bad Request status and the error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse the board ID from the URL parameter.
	boardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// If parsing the ID fails, respond with a 400 Bad Request status and an invalid ID error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Attempt to update the board in the database using the provided data and ID.
	success, err := model.PutBoard(json, boardID)
	if success {
		// If the update is successful, respond with a 200 OK status and a success message.
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		// If the update fails, respond with a 400 Bad Request status and the error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

// DeleteBoardByID handles the deletion of a board by its ID.
//
// It expects the following:
// - The board ID as a URL parameter.
//
// The function returns a JSON response with a status and a message indicating the result.
// func DeleteBoardByID(c *gin.Context) {
// 	// Parse the board ID from the URL parameter.
// 	boardID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		// If parsing the ID fails, respond with a 400 Bad Request status and an invalid ID error message.
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
// 		return
// 	}

// 	// Attempt to delete the board from the database using the provided ID.
// 	success, err := model.DeleteBoard(boardID)
// 	if success {
// 		// If the deletion is successful, respond with a 200 OK status and a success message.
// 		c.JSON(http.StatusOK, gin.H{"message": "Board deleted successfully"})
// 	} else {
// 		// If the deletion fails, respond with a 400 Bad Request status and the error message.
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 	}
// }
