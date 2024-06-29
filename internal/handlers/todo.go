package handlers

import (
	"net/http"

	"github.com/FlawlessByte/todo-app/internal/models"
	"github.com/FlawlessByte/todo-app/internal/repository"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Repo repository.TodoRepository
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos := h.Repo.GetAll()
	c.IndentedJSON(http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo

	// Call BindJSON to bind the received JSON to
	// todo.
	if err := c.BindJSON(&todo); err != nil {
		return
	}
	// Add the new todo to the slice.
	h.Repo.Create(todo)
	c.IndentedJSON(http.StatusCreated, todo)
}
