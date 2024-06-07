// Package handlers manages HTTP requests.
package handlers

import (
	"net/http"
	"strconv"

	"example.com/kanban/models"
	"github.com/gin-gonic/gin"
)

// DeleteBoardByID handles the deletion of a board by its ID.
//
// It expects the following:
// - The board ID as a URL parameter.
//
// The function returns a JSON response with a status and a message indicating the result.
func DeleteBoardByID(c *gin.Context) {
	// Parse the board ID from the URL parameter.
	boardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// If parsing the ID fails, respond with a 400 Bad Request status and an invalid ID error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Attempt to delete the board from the database using the provided ID.
	success, err := models.DeleteBoard(boardID)
	if success {
		// If the deletion is successful, respond with a 200 OK status and a success message.
		c.JSON(http.StatusOK, gin.H{"message": "Board deleted successfully"})
	} else {
		// If the deletion fails, respond with a 400 Bad Request status and the error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
