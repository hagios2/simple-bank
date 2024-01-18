package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hagios2/simple-bank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server and sets up routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//setup routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.createAccount)

	server.router = router
	return server
}

func (server *Server) STart(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
