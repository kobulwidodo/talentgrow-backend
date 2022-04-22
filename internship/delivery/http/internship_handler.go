package http

import (
	"net/http"
	"talentgrow-backend/domain"
	"talentgrow-backend/middleware"
	"talentgrow-backend/utils"

	"github.com/gin-gonic/gin"
)

type InternshipHandler struct {
	InternshipUseCase domain.InternshipUseCase
}

func NewInternshipHandler(r *gin.RouterGroup, iu domain.InternshipUseCase, jwtMiddleware gin.HandlerFunc, mustAdmin gin.HandlerFunc) {
	handler := &InternshipHandler{InternshipUseCase: iu}
	api := r.Group("/internship")
	{
		api.POST("/", jwtMiddleware, mustAdmin, handler.CreateInternship)
		api.GET("/:id", handler.GetInternshipById)
		api.GET("/", handler.GetInternships)
		api.PUT("/:id", jwtMiddleware, mustAdmin, handler.UpdateInternship)
		api.DELETE("/:id", jwtMiddleware, mustAdmin, handler.DeleteInternship)
	}
}

func (h *InternshipHandler) CreateInternship(c *gin.Context) {
	input := new(domain.CreateInternship)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	input.UserId = c.MustGet("auth").(middleware.CustomClaim).Id
	if err := h.InternshipUseCase.CreateInternship(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully created internship", nil))
}

func (h *InternshipHandler) GetInternshipById(c *gin.Context) {
	input := new(domain.FindInternship)
	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	internship, err := h.InternshipUseCase.GetInternshipById(input)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully fetch data", internship))
}

func (h *InternshipHandler) GetInternships(c *gin.Context) {
	internship, err := h.InternshipUseCase.GetInternships()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully fetch data", internship))
}

func (h *InternshipHandler) UpdateInternship(c *gin.Context) {
	input := new(domain.UpdateInternship)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	inputUri := new(domain.FindInternship)
	if err := c.ShouldBindUri(&inputUri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	input.Id = inputUri.Id
	if err := h.InternshipUseCase.UpdateInternship(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully updated date", nil))
}

func (h *InternshipHandler) DeleteInternship(c *gin.Context) {
	input := new(domain.FindInternship)
	if err := c.ShouldBindUri(input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	if err := h.InternshipUseCase.DeleteInternship(input.Id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully deleted data", nil))
}
