package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"simple-rest/types"
	"simple-rest/utils/logger"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	logger *logger.Logger
	psql   *sql.DB
}

func NewAuthRouter(logger *logger.Logger, psql *sql.DB) *AuthRouter {
	return &AuthRouter{
		logger: logger,
		psql:   psql,
	}
}

func (r *AuthRouter) register(c *gin.Context) {
	var user types.User

	if err := c.BindJSON(&user); err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect json data",
		})
		return
	}

	_, err := r.psql.Exec("INSERT INTO authors (name, last_name, login, password) VALUES ($1, $2, $3, $4)", user.Name, user.LastName, user.Login, user.Password)
	if err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't insert this data",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}

func (r *AuthRouter) login(c *gin.Context) {
	var user types.User

	if err := c.BindJSON(&user); err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect json data",
		})
		return
	}

	stmt, err := r.psql.Prepare("SELECT id FROM authors WHERE login=$1 AND password=$2")
	if err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find user",
		})
		return
	}

	var id int64
	if err := stmt.QueryRow(user.Login, user.Password).Scan(&id); err != nil {
		r.logger.PrintWarning(fmt.Sprintf("%v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't find user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user found",
	})
}
