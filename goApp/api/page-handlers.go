package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/EmilioCliff/car-rental/templates"
	"github.com/gin-gonic/gin"
)

func(server *Server) homePage(ctx *gin.Context) {
	indexBody := templates.HomeBody()
	if err := templates.HomePage(indexBody).Render(ctx, ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error generating template: %v", err) })
	}
}

func (server *Server) safariPage(ctx *gin.Context) {
	tmpl := templates.SafariBody()
	if err := templates.SafariPage(tmpl).Render(ctx, ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error generating template: %v", err) })
	}
}

func(server *Server) safariNumPage(ctx *gin.Context) {
	safariIDString := ctx.Param("id")
	safariIDInt, err := strconv.Atoi(safariIDString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error converting safari id to int: %v", err) })
	}
	safari := templates.GetSafari(int32(safariIDInt))
	tmpl := templates.SafariNumBody(safari)
	if err := templates.SafariNumPage(tmpl).Render(ctx, ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error generating template: %v", err) })
	}
}

func (server *Server) getTaxiPage(ctx *gin.Context) {
	tmpl := templates.TaxiBody()
	if err := templates.TaxiPage(tmpl).Render(ctx, ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error generating template: %v", err) })
	}
}