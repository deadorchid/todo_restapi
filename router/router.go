package router

import (
	"database/sql"
	"simple-rest/utils/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(logger *logger.Logger, db *sql.DB) *gin.Engine {
	r := gin.Default()

	todoRouter := NewTodoRouter(logger, db)

	r.GET("/todo", todoRouter.getTodos)
	r.POST("/todo", todoRouter.createTodo)
	r.GET("/todo/:id", todoRouter.todoId)
	r.DELETE("/todo/:id", todoRouter.deleteTodo)
	r.PUT("/todo/:id", todoRouter.updateTodo)

	return r
}
