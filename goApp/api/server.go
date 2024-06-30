package api

import (
	// "html/template"

	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	var server Server

	server.setRoutes()

	return &server
}

func (server *Server) setRoutes() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/static/") {
			c.Header("Cache-Control", "public, max-age=2592000")
		}
		c.Next()
	})

	router.Use(static.Serve("/static", static.LocalFile("./static", false)))

	router.GET("/", server.homePage)
	router.GET("/safaris", server.safariPage)
	router.GET("/safaris/:id", server.safariNumPage)
	router.GET("/get-taxi", server.getTaxiPage)
	router.GET("/car-hire", server.carHire)

	server.router = router
}

func (server *Server)Start(address string) error {
	return server.router.Run()
}

// func (server *Server) render(ctx *gin.Context, status int, template templ.Component) error {
// 	ctx.Status(status)
// 	return template.Render(ctx.Request.Context(), ctx.Writer)
// }