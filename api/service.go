package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func DeclareServer() *Server {
	e := gin.New()
	server := &Server{
		Engine: e,
	}
	// Attach middleware
	e.Use(server.middlewareCheck())
	server.Engine = e
	return server
}

// Middleware
// Chi su dung duoc trong file nay
// Ben ngoai goi vao khong duoc
func (s *Server) middlewareCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
			return
		}
		ctx.Next()
	}
}

// Run Server
// Su dung duoc o ben ngoai server
func (s *Server) Run() {
	s.Engine.Run(":8000")
}
