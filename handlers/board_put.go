// Package handlers manages HTTP requests.
package handlers

import (
	"net/http"
	"strconv"

	"example.com/kanban/models"
	"github.com/gin-gonic/gin"
)

/* 
PutBoard handles the update of a board.

It expects the following:
- A JSON body containing the updated fields for the Board struct.
- The board ID as a URL parameter.

The function returns a JSON response with the status and a message indicating the result. 
*/
func PutBoard(c *gin.Context) {
	// Initialize a variable to hold the JSON body parsed into a Board struct.
	var json models.Board

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
	success, err := models.PutBoard(json, boardID)
	if success {
		// If the update is successful, respond with a 200 OK status and a success message.
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		// If the update fails, respond with a 400 Bad Request status and the error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
