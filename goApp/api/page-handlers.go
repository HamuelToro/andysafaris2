package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/EmilioCliff/car-rental/templates"
	"github.com/a-h/templ"
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

func (server *Server) carHire(ctx *gin.Context) {
	tmpl := templates.CarHireBody()
	if err := templates.CarHirePage(tmpl).Render(ctx, ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error generating template: %v", err) })
	}
}

type safariRequestParams struct {
	ID int32 `json:"id" binding:"required"`
}

type carHireRequestParams struct {
	ID int16 `json:"id" binding:"required"`
}

// work on nairobi-tour
// type nairobiTourRequestParams struct {
	
// }

// displays the contact form
func (server *Server) contactUs(ctx *gin.Context) {
	service := ctx.Param("service")
	var body templ.Component
	switch (service) {
	case "safari":
		var req safariRequestParams
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
			return
		}
		
		log.Println(req.ID)
		safari := templates.GetSafari(req.ID)
		body = templates.SafariForm(safari)
	case "car-hire":
		var req carHireRequestParams
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
			return
		}

		vehicle := templates.GetVehicle(req.ID)
		body = templates.CarHireForm(vehicle)
	case "nairobi-tour":
	default:
		log.Println("bad service type: ", service)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("service type isnt recoqnized: %s", service)})
		return
	}

	if err := templates.ContactUsPage(body).Render(ctx, ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}
}

func (server *Server) contactUsPage(ctx *gin.Context) {
	tmpl := templates.EnquiryForm()
	if err := templates.ContactUsPage(tmpl).Render(ctx, ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error generating template: %v", err) })
	}
}

func (server *Server) safariContact(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	safari := templates.GetSafari(int32(idInt))
	tmpl := templates.SafariForm(safari)
	if err := templates.ContactUsPage(tmpl).Render(ctx, ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error generating template: %v", err) })
	}
}

func (server *Server) carHireContact(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	vehicle := templates.GetVehicle(int16(idInt))
	tmpl := templates.CarHireForm(vehicle)
	if err := templates.ContactUsPage(tmpl).Render(ctx, ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error generating template: %v", err) })
	}
}
