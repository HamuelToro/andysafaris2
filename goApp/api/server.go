package api

import (
	"html/template"
	"log"
	"net/http"

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

	// router.Use(func(c *gin.Context) {
	// 	if strings.HasPrefix(c.Request.URL.Path, "/static/") {
	// 		c.Header("Cache-Control", "public, max-age=2592000")
	// 	}
	// 	c.Next()
	// })

	// router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/**")
	router.StaticFS("/static", http.Dir("./static"))
		// router.GET("/", server.homePage)
	router.GET("/", func(ctx *gin.Context) {
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			log.Println(err)
			return
		}

		err = tmpl.Execute(ctx.Writer, []string{"hey", "you", "there", "mathafaka"})
		if err != nil {
			log.Println(err)
			return
		}
		// ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Posts"})
	})
	server.router = router
}

func (server *Server)Start(address string) error {
	return server.router.Run()
}

// func (server *Server) render(ctx *gin.Context, status int, template templ.Component) error {
// 	ctx.Status(status)
// 	return template.Render(ctx.Request.Context(), ctx.Writer)
// }