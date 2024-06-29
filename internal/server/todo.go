package server

import (
	"github.com/FlawlessByte/todo-app/internal/handlers"
	"github.com/FlawlessByte/todo-app/internal/repository"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func GinServer() *Server {
	repo := repository.NewTodoRepository()
	handler := &handlers.TodoHandler{Repo: *repo}

	router := gin.Default()
	api := router.Group("/v1")
	{
		api.GET("/todos", handler.GetTodos)
		api.POST("/todos", handler.CreateTodo)
	}

	return &Server{
		Router: router,
	}
}
