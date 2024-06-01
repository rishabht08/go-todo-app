package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishabht08/go-todo-app/internal/router"
)

func routes(db *sql.DB) http.Handler {
	routeHandler := gin.Default()

	routeGroup := routeHandler.Group("/api/v1")

	router.MakeRoutes(db, routeGroup)

	return routeHandler
}
