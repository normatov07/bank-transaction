package api

import (
	db "bank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//Server to HTTP request to our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

//NewServer creates new HTTH server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

//Start run HTTP server on a specific address
func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

//Render error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
