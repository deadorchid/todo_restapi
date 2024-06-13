package router

import (
	"simple-rest/types"
	"simple-rest/utils/logger"

	"github.com/gin-gonic/gin"
)

var todos []types.Todo = []types.Todo{
	{
		Title:  "Приготовить обэд",
		Status: 0,
		Author: types.Person{
			Name:     "Daniil",
			LastName: "Mikhaylov",
		},
	},
	{
		Title:  "Приготовить завтрак",
		Status: 0,
		Author: types.Person{
			Name:     "Daniil",
			LastName: "Mikhaylov",
		},
	},
	{
		Title:  "Поспать",
		Status: 1,
		Author: types.Person{
			Name:     "Daniil",
			LastName: "Mikhaylov",
		},
	},
}

func NewRouter(logger *logger.Logger) *gin.Engine {
	r := gin.Default()

	todoRouter := NewTodoRouter(logger)

	r.GET("/todo", todoRouter.getTodos)
	r.GET("/todo/:id", todoRouter.todoId)
	r.POST("/todo", todoRouter.createTodo)
	r.DELETE("/todo/:id", todoRouter.deleteTodo)
	r.PUT("/todo/:id", todoRouter.updateTodo)

	return r
}
