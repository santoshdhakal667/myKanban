// Package handlers manages HTTP requests.
package handlers

import (
	"log"
	"net/http"

	"example.com/kanban/models"
	"github.com/gin-gonic/gin"
)

// PostBoard handles the creation of a new board.
//
// It expects the following:
// - A JSON body containing the fields for the new Board struct.
//
// The function returns a JSON response with the status and a message indicating the result.
func PostBoard(c *gin.Context) {
	// Initialize a variable to hold the JSON body parsed into a Board struct.
	var json models.Board

	// Parse the JSON request body into the Board struct.
	if err := c.ShouldBindJSON(&json); err != nil {
		// If parsing fails, respond with a 400 Bad Request status and the error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// Log the error for server-side debugging.
		log.Println(err.Error())
		return
	}

	// Log the successfully parsed Board struct for debugging purposes.
	log.Println(json)

	// Attempt to create the new board in the database using the provided data.
	success, err := models.PostBoard(json)
	if success {
		// If the creation is successful, respond with a 200 OK status and a success message.
		c.JSON(http.StatusOK, gin.H{"success": success})
	} else {
		// If the creation fails, respond with a 400 Bad Request status and the error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// Log the error for server-side debugging.
		log.Println(err)
	}
}
