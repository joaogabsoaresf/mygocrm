package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/go-stock-watcher/schemas"
)

func UpdateLeadHandler(ctx *gin.Context) {
	request := UpdateLeadRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	lead := schemas.Lead{}

	if err := db.First(&lead, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("lead with id: %s not found", id))
		return
	}

	//update lead
	if request.Name != "" {
		lead.Name = request.Name
	}
	if request.Email != "" {
		lead.Email = request.Email
	}
	if request.Phone != "" {
		lead.Phone = request.Phone
	}
	if request.Client != nil {
		lead.Client = *request.Client
	}
	if request.Age > 0 {
		lead.Age = request.Age
	}

	// save lead
	if err := db.Save(&lead).Error; err != nil {
		logger.Errorf("error updating lead: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("erro updating lead id: %s", id))
		return
	}
	sendSuccess(ctx, "update", lead)
}
