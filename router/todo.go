package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"simple-rest/types"
	"simple-rest/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoRouter struct {
	logger *logger.Logger
	psql   *sql.DB
}

func NewTodoRouter(logger *logger.Logger, psql *sql.DB) *TodoRouter {
	return &TodoRouter{
		logger: logger,
		psql:   psql,
	}
}

func (r *TodoRouter) getTodos(c *gin.Context) {
	var res types.Todo
	var author types.Person
	var todos []types.Todo

	rows, err := r.psql.Query("SELECT * FROM todos")

	if err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something gone wrong",
		})
		return
	}

	for rows.Next() {
		rows.Scan(&res.Id, &author.Id, &res.Title, &res.Status)
		row, err := r.psql.Query(fmt.Sprintf("SELECT * FROM authors WHERE id = %d", author.Id))
		if err != nil {
			r.logger.PrintWarning(fmt.Sprintf("%v", err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something gone wrong",
			})
			return
		}

		for row.Next() {
			row.Scan(&author.Id, &author.Name, &author.LastName)
		}

		row.Close()
		res.Author = author
		todos = append(todos, res)
	}

	rows.Close()

	c.JSON(http.StatusOK, todos)
}

func (r *TodoRouter) todoId(c *gin.Context) {
	var todo types.Todo
	var author types.Person

	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert id",
		})
		return
	}

	rows, err := r.psql.Query(fmt.Sprintf("SELECT * FROM todos WHERE id=%d", id))
	if err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something gone wrong",
		})
		return
	}

	for rows.Next() {
		rows.Scan(&todo.Id, &author.Id, &todo.Title, &todo.Status)
		row, err := r.psql.Query(fmt.Sprintf("SELECT * FROM authors WHERE id = %d", author.Id))
		if err != nil {
			r.logger.PrintWarning(fmt.Sprintf("%v", err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something gone wrong",
			})
			return
		}

		for row.Next() {
			row.Scan(&author.Id, &author.Name, &author.LastName)
		}

		row.Close()
	}

	todo.Author = author

	if todo.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(http.StatusFound, todo)
}

func (r *TodoRouter) createTodo(c *gin.Context) {
	var req types.Todo
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	stmt, err := r.psql.Prepare("INSERT INTO todos (author_id, title, status) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't insert this todo",
		})
		return
	}

	var id int64

	err = stmt.QueryRow(req.Author.Id, req.Title, req.Status).Scan(&id)
	if err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't insert get id of this todo",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("todo created with id: %d", id),
	})
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

	_, err = r.psql.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to delete",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo deleted",
	})
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

	var req types.Todo

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	stmt, err := r.psql.Prepare("UPDATE todos SET author_id=$1, title=$2, status=$3 WHERE id=$4 RETURNING id")
	if err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't insert this todo",
		})
		return
	}

	err = stmt.QueryRow(req.Author.Id, req.Title, req.Status, id).Scan(&id)
	if err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't update this todo",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("todo with id: %d changed", id),
	})
}
