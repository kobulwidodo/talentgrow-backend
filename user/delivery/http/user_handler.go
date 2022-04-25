package http

import (
	"net/http"
	"talentgrow-backend/domain"
	"talentgrow-backend/middleware"
	"talentgrow-backend/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

func NewUserHandler(r *gin.RouterGroup, uu domain.UserUseCase, jwtMiddleware gin.HandlerFunc) {
	handler := &UserHandler{UserUseCase: uu}
	api := r.Group("/auth")
	{
		api.POST("/signup", handler.SignUp)
		api.POST("/signin", handler.SignIn)
	}
	r.GET("/me", jwtMiddleware, handler.GetMe)
}

func (h *UserHandler) SignUp(c *gin.Context) {
	input := new(domain.UserSignUp)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	if err := h.UserUseCase.SignUp(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("user successfully registered", nil))
}

func (h *UserHandler) SignIn(c *gin.Context) {
	input := new(domain.UserSignIn)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	token, err := h.UserUseCase.SignIn(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully login", map[string]string{"token": token}))
}

func (h *UserHandler) GetMe(c *gin.Context) {
	email := c.MustGet("auth").(middleware.CustomClaim).Email
	user, err := h.UserUseCase.GetMe(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully get me", user))
	return
}
