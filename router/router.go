package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type Todo struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Author Person `json:"author"`
}

var todos []Todo = []Todo{
	{
		"Приготовить обэд",
		0,
		Person{
			"Daniil",
			"Mikhaylov",
		},
	},
	{
		"Приготовить завтрак",
		0,
		Person{
			"Daniil",
			"Mikhaylov",
		},
	},
	{
		"Поспать",
		1,
		Person{
			"Daniil",
			"Mikhaylov",
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
		var requestBody Todo
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

	return r
}
