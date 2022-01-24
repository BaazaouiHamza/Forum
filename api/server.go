package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/hamza-baazaoui/forum/db/sqlc"
	"github.com/hamza-baazaoui/forum/token"
	"github.com/hamza-baazaoui/forum/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, db db.Store) (*Server, error) {
	tokenMaker, err := token.NewJwt(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      db,
		tokenMaker: tokenMaker,
	}
	server.setUpRouter()
	return server, nil
}
func (server *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/posts", server.createPost)
	router.PATCH("/posts/:id", server.updatePost)
	router.GET("/posts", server.listPost)
	router.GET("/posts/:id", server.getPostById)
	router.DELETE("/posts/:id", server.deletePost)

	server.router = router
}

//Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
