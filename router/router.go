package router

import (
	"fmt"
	"net/http"
	"simple-rest/types"
	"strconv"

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

	r.GET("/todo", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	r.GET("/todo/:id", func(c *gin.Context) {
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
	})

	r.POST("/todo", func(c *gin.Context) {
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
	})

	r.DELETE("/todo/:id", func(c *gin.Context) {
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
	})

	r.PUT("/todo/:id", func(c *gin.Context) {
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
	})

	return r
}
