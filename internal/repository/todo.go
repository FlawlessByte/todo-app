package repository

import "github.com/FlawlessByte/todo-app/internal/models"

type TodoRepository struct {
	todos []models.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos: make([]models.Todo, 0),
	}
}

func (r *TodoRepository) GetAll() []models.Todo {
	return r.todos
}

func (r *TodoRepository) Create(todo models.Todo) {
	r.todos = append(r.todos, todo)
}
