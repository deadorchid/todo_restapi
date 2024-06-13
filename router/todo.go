package router

import (
	"fmt"
	"net/http"
	"simple-rest/types"
	"simple-rest/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoRouter struct {
	logger *logger.Logger
}

func NewTodoRouter(logger *logger.Logger) *TodoRouter {
	return &TodoRouter{
		logger: logger,
	}
}

func (r *TodoRouter) getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func (r *TodoRouter) todoId(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert id",
		})
		return
	}

	if id > len(todos)-1 || id < 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to found given id",
		})
		return
	}

	c.JSON(http.StatusFound, todos[id])
}

func (r *TodoRouter) createTodo(c *gin.Context) {
	var requestBody types.Todo
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	todos = append(todos, requestBody)
	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("todo created with id: %d", len(todos)-1),
	})
	r.logger.PrintInfo("new todo created")
}

func (r *TodoRouter) deleteTodo(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert id",
		})
		return
	}

	if id > len(todos)-1 || id < 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to found given id",
		})
		return
	}

	todos = append(todos[:id], todos[id+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"message": "todo deleted",
	})
	r.logger.PrintInfo("todo deleted")
}

func (r *TodoRouter) updateTodo(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert id",
		})
		return
	}

	if id > len(todos)-1 || id < 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to found given id",
		})
		return
	}

	var requestBody types.Todo

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	todos[id] = requestBody
	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("todo with id: %d changed", id),
	})
	r.logger.PrintInfo("todo updated")
}
