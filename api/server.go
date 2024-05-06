package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/hagios2/simple-bank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server and sets up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//registering validators here, for now its just currency validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			return nil
		}
	}

	//setup routes
	router.POST("/users", server.createUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.getAccountList)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
