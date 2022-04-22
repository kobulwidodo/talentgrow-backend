package http

import (
	"net/http"
	"talentgrow-backend/domain"
	"talentgrow-backend/middleware"
	"talentgrow-backend/utils"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	EventUseCase domain.EventUseCase
}

func NewEventHandler(r *gin.RouterGroup, eu domain.EventUseCase, jwtMiddleware gin.HandlerFunc, mustAdmin gin.HandlerFunc) {
	handler := &EventHandler{EventUseCase: eu}
	api := r.Group("/event")
	{
		api.POST("/", jwtMiddleware, mustAdmin, handler.CreateEvent)
		api.GET("/", handler.GetAllEvent)
		api.GET("/type/:type", handler.GetAllEventByType)
		api.GET("/:id", handler.GetEventById)
		api.PUT("/:id", handler.UpdateEvent)
		api.DELETE("/:id", handler.DeleteEvent)
	}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	input := new(domain.CreateEventDto)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	input.UserId = c.MustGet("auth").(middleware.CustomClaim).Id
	if err := h.EventUseCase.Create(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully created new event", nil))
}

func (h *EventHandler) GetAllEvent(c *gin.Context) {
	events, err := h.EventUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get all events", events))
}

func (h *EventHandler) GetAllEventByType(c *gin.Context) {
	uri := new(domain.FindEventsType)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	events, err := h.EventUseCase.GetByType(uri.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get events", events))
}

func (h *EventHandler) GetEventById(c *gin.Context) {
	uri := new(domain.FindEventUri)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	event, err := h.EventUseCase.GetById(uri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get event", event))
}

func (h *EventHandler) UpdateEvent(c *gin.Context) {
	input := new(domain.UpdateEventDto)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	uri := new(domain.FindEventUri)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	input.Id = uri.Id
	if err := h.EventUseCase.Update(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully updated event", nil))
}

func (h *EventHandler) DeleteEvent(c *gin.Context) {
	uri := new(domain.FindEventUri)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	if err := h.EventUseCase.Delete(uri.Id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully deleted event", nil))
}