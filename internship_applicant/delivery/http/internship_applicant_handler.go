package http

import (
	"net/http"
	"talentgrow-backend/domain"
	"talentgrow-backend/middleware"
	"talentgrow-backend/utils"

	"github.com/gin-gonic/gin"
)

type InternshipApplicantHandler struct {
	internshipApplicantUseCase domain.InternshipAppilcantUseCase
}

func NewInternshipApplicantHandler(r *gin.RouterGroup, iau domain.InternshipAppilcantUseCase, jwtMiddleware gin.HandlerFunc) {
	handler := &InternshipApplicantHandler{internshipApplicantUseCase: iau}
	api := r.Group("/internship-applicant")
	{
		api.POST("/create/:internship_id", jwtMiddleware, handler.ApplyInternship)
	}
}

func (h *InternshipApplicantHandler) ApplyInternship(c *gin.Context) {
	input := new(domain.ApplyInternship)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	inputUri := new(domain.FindInternshipUri)
	if err := c.ShouldBindUri(inputUri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	input.UserId = c.MustGet("auth").(middleware.CustomClaim).Id
	input.InternshipId = inputUri.InternshipId
	if err := h.internshipApplicantUseCase.Apply(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully apply", nil))
}
