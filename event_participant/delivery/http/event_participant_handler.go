package http

import (
	"net/http"
	"talentgrow-backend/domain"
	"talentgrow-backend/middleware"
	"talentgrow-backend/utils"

	"github.com/gin-gonic/gin"
)

type EventParticipantHandler struct {
	eventParticipantUsecase domain.EventParticipantUsecase
}

func NewEventParticipantHandler(r *gin.RouterGroup, epu domain.EventParticipantUsecase, jwtMiddleware gin.HandlerFunc) {
	handler := &EventParticipantHandler{eventParticipantUsecase: epu}
	api := r.Group("/event-participant")
	{
		api.POST("/:event_id", jwtMiddleware, handler.RegisterEvent)
		api.GET("/status/:event_id", jwtMiddleware, handler.CheckIsRegistered)
	}
}

func (h *EventParticipantHandler) RegisterEvent(c *gin.Context) {
	input := new(domain.CreateEventParticipant)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	uri := new(domain.FindEventParticipantUri)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	input.UserId = c.MustGet("auth").(middleware.CustomClaim).Id
	input.EventId = uri.EventId
	if err := h.eventParticipantUsecase.Create(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully register this event", nil))
}

func (h *EventParticipantHandler) CheckIsRegistered(c *gin.Context) {
	uri := new(domain.FindEventParticipantUri)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	userId := c.MustGet("auth").(middleware.CustomClaim).Id
	data, err := h.eventParticipantUsecase.CheckIsRegisterd(userId, uri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get status", data))
}