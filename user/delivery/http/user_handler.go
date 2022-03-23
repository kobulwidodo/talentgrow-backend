package http

import (
	"talentgrow-backend/domain"
	"talentgrow-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

func NewUserHandler(r *gin.RouterGroup, uu domain.UserUseCase) {
	handler := &UserHandler{UserUseCase: uu}
	api := r.Group("/auth")
	{
		api.POST("/signup", handler.SignUp)
		api.POST("/signin", handler.SignIn)
	}
}

func (u *UserHandler) SignUp(c *gin.Context) {
	input := new(domain.UserSignUp)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	if err := u.UserUseCase.SignUp(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("user successfully registered", nil))
}

func (u *UserHandler) SignIn(c *gin.Context) {
	input := new(domain.UserSignIn)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	token, err := u.UserUseCase.SignIn(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully login", map[string]string{"token": token}))
}
