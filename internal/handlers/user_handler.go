package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishabht08/go-todo-app/internal/repositories"
)

type Repository struct {
	Db *repositories.PostgresRepo
}

var Repo *Repository

func NewUserRepo(db *sql.DB, route *gin.RouterGroup) {
	Repo = &Repository{
		Db: repositories.NewRepo(db),
	}

	route.POST("/dummy", Repo.CreateDummyUserHandler)
}

func (m *Repository) CreateDummyUserHandler(c *gin.Context) {

	err := Repo.Db.CreateDummyUser()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
	})

}
