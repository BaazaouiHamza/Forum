package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hamza-baazaoui/forum/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(db db.Store) *Server {
	server := &Server{store: db}
	router := gin.Default()

	router.POST("/users", server.createUser)

	router.POST("/posts", server.createPost)
	router.PATCH("/posts/:id", server.updatePost)
	router.GET("/posts", server.listPost)
	router.GET("/posts/:id", server.getPostById)
	router.DELETE("/posts/:id", server.deletePost)

	server.router = router
	return server
}

//Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
