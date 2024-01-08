package Controllers

import (
	"TODO/Config"
	"TODO/Models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// Get all TODO Data
func GetTODO(c echo.Context) error {

	// Get database
	db := Config.GetDB()

	// Get TODO elements from database
	var todo []*Models.TODO
	if err := db.Find(&todo); err.Error != nil {
		data := map[string]interface{}{
			"message": err.Error.Error(),
		}
		return c.JSON(http.StatusOK, data)
	}

	// Build response
	response := map[string]interface{}{
		"message": "Get: Success!",
		"data":    todo,
	}

	// Send response
	return c.JSON(http.StatusOK, response)
}

// Create TODO Data
func CreateTODO(c echo.Context) error {

	// Get database
	db := Config.GetDB()

	// Bind data
	todo := new(Models.TODO)
	if err := c.Bind(todo); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	// Validate the TODO struct
	validate := validator.New()
	if err := validate.Struct(todo); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	// Create new TODO element
	if err := db.Create(&todo).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// Build response
	response := map[string]interface{}{
		"message": "Creation: Success!",
		"data":    todo,
	}

	// Send response
	return c.JSON(http.StatusOK, response)
}

// Update TODO Data
func UpdateTODO(c echo.Context) error {

	// Get database
	db := Config.GetDB()

	// Bind data
	todo := new(Models.TODO)
	if err := c.Bind(todo); err != nil {
		println("bind error")
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// Validate the TODO struct
	validate := validator.New()
	if err := validate.Struct(todo); err != nil {
		println("validator error")
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	// Retrieve ID from the URL parameter
	id := c.Param("id")

	// Retrieve existing TODO element from db
	existingTodo := new(Models.TODO)
	if err := db.First(&existingTodo, id).Error; err != nil {
		println("first error")
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// Set new TODO fields
	existingTodo.Title = todo.Title
	existingTodo.Deadline = todo.Deadline
	existingTodo.Done = todo.Done

	// Save new TODO fields
	if err := db.Save(&existingTodo).Error; err != nil {
		println("save error")
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// Build response
	response := map[string]interface{}{
		"message": "Update: Success!",
		"data":    existingTodo,
	}

	// Send response
	return c.JSON(http.StatusOK, response)
}

// Delete TODO Data
func DeleteTODO(c echo.Context) error {

	// Get database
	db := Config.GetDB()

	// Retrieve ID from the URL parameter
	id := c.Param("id")

	// Delete TODO element from database
	todo := new(Models.TODO)
	err := db.Delete(&todo, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	// Build response
	response := map[string]interface{}{
		"message": "Deletion: Success!",
	}

	// Send response
	return c.JSON(http.StatusOK, response)
}
