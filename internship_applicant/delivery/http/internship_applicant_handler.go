package http

import (
	"fmt"
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
		api.GET("/:id", jwtMiddleware, handler.FindOne)
		api.POST("/create/cv/:id", jwtMiddleware, handler.UploadCv)
		api.GET("/status/:internship_id", jwtMiddleware, handler.CheckRegistered)
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
	userId := c.MustGet("auth").(middleware.CustomClaim).Id
	input.UserId = userId
	input.InternshipId = inputUri.InternshipId
	id, err := h.internshipApplicantUseCase.Apply(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully apply", map[string]uint{"id": id}))
}

func (h *InternshipApplicantHandler) FindOne(c *gin.Context) {
	input := new(domain.FindApplicant)
	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	data, err := h.internshipApplicantUseCase.FindOne(input)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get data", data))
}

func (h *InternshipApplicantHandler) CheckRegistered(c *gin.Context) {
	uri := new(domain.FindInternshipUri)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	userId := c.MustGet("auth").(middleware.CustomClaim).Id
	data, err := h.internshipApplicantUseCase.CheckIsRegistered(userId, uri.InternshipId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get status", data))
}

func (h *InternshipApplicantHandler) UploadCv(c *gin.Context) {
	uri := new(domain.FindApplicant)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	userId := c.MustGet("auth").(middleware.CustomClaim).Id
	path := fmt.Sprintf("cv/%d-%s", userId, file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	if err := h.internshipApplicantUseCase.UploadCv(path, uri.Id); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully upload cv", nil))
}
