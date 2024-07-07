package api

import (
	// "html/template"

	"strings"

	"github.com/EmilioCliff/car-rental/utils"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	sender utils.GmailSender
}

func NewServer(sender utils.GmailSender) *Server {
	var server Server

	server.setRoutes()
	server.sender = sender

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
	router.GET("/nairobi-tour", server.nairobiTourPage)
	router.GET("/get-taxi", server.getTaxiPage)
	router.GET("/car-hire", server.carHire)
	router.GET("/safari/:id/contact-form", server.safariContact)
	router.GET("/car-hire/:id/contact-form", server.carHireContact)
	router.GET("/nairobi-tour/contact-form", server.nairobiTourContact)
	router.GET("/contact-us", server.contactUsPage)
	router.POST("/contact-form/:form", server.formHandlers)

	server.router = router
}

func (server *Server)Start(address string) error {
	return server.router.Run()
}

func (server *Server) errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}

// func (server *Server) render(ctx *gin.Context, status int, template templ.Component) error {
// 	ctx.Status(status)
// 	return template.Render(ctx.Request.Context(), ctx.Writer)
// }