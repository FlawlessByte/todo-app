package main

import (
	"github.com/FlawlessByte/todo-app/internal/server"
)

func main() {
	srv := server.GinServer()
	srv.Router.Run("localhost:8080")
}
