package router

import (
	"simple-rest/types"

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

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/todo", todosRouter)
	r.GET("/todo/:id", todoIdRouter)
	r.POST("/todo", createTodoRouter)
	r.DELETE("/todo/:id", deleteTodoRouter)
	r.PUT("/todo/:id", updateTodoRouter)

	return r
}
