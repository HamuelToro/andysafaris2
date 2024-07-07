package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/EmilioCliff/car-rental/templates"
	"github.com/gin-gonic/gin"
)

// type enqueryForm struct {
// 	fullName string
// 	email string
// 	phoneNumber string
// 	serviceEnguery string
// 	additionalMessage string
// }

// type safariForm struct {
// 	firstName string
// 	lastName string
// 	email string
// 	phoneNumber string
// 	safariDate string
// 	arrivalDate string
// 	noOfAdults int16
// 	noOfChildren int16
// 	additionalMessage string
// }

// type carHireForm struct {
// 	firstName string
// 	lastName string
// 	email string
// 	phoneNumber string
// 	duration string
// 	pickupDay string
// 	additionalMessage string
// }

// type reviewForm struct {
// 	fullName string
// 	email string
// 	review string
// }

func (server *Server) formHandlers(ctx *gin.Context) {
	form := ctx.Param("form")

	switch (form) {
	case "enquery":
	case "safari":
		var req templates.SafariFormPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
			return
		}

		fmt.Println(req)
		t := templates.SafariFormComponent(req)

		var buf bytes.Buffer
		if err := t.Render(ctx, &buf); err != nil {
			ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
			return
		}

		emailContent := buf.String()
		fmt.Println(emailContent) // For debugging

		// if err := server.sender.SendMAil("Safari Booking Confirmation", emailContent, "text/plain", []string{"emiliocliff@gmail.com"}, nil, nil, nil, nil); err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		// 	return
		// }

	// 
	case "nairobi-tour":
	case "car-hire":
	case "review":
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unknown form"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "Enquiry sent successfull"})
}