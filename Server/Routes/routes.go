package Routes

import (
	"TODO/Controllers"

	"github.com/labstack/echo/v4"
)

// Setup for routing
func SetupRoutes(e *echo.Echo) {
	// Get TODO data
	e.GET("/todo", Controllers.GetTODO)
	// Create new TODO
	e.POST("/todo", Controllers.CreateTODO)
	// Modify a TODO, given its id
	e.PUT("/todo/:id", Controllers.UpdateTODO)
	// Delete a TODO, given its id
	e.DELETE("/todo/:id", Controllers.DeleteTODO)
}
