package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rishabht08/go-todo-app/internal/handlers"
)

func MakeRoutes(db *sql.DB, route *gin.RouterGroup) {
	handlers.NewUserRepo(db, route)
}
