// Package handlers manages HTTP requests.
package handlers

import (
	"log"
	"net/http"

	"example.com/kanban/models"
	"github.com/gin-gonic/gin"
)

// GetBoard handles retrieving a list of boards.
//
// It retrieves a fixed number of boards (20 in this case) from the database.
// The function returns a JSON response with the boards data or an error message.
func GetBoard(c *gin.Context) {
	// Retrieve a list of boards from the database.
	boards, err := models.GetBoard(20)
	if err != nil {
		// If retrieval fails, log the error and terminate the application.
		log.Fatal(err)
	}

	// If no boards are found, respond with a 404 Not Found status and a message.
	if boards == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No record found"})
		return
	}

	// If boards are found, respond with a 200 OK status and the boards data.
	c.JSON(http.StatusOK, gin.H{"data": boards})
}

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
	board, err := models.GetBoardByID(id)
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
