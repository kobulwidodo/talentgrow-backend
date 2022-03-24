package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"talentgrow-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CustomClaim struct {
	Id      uint
	IsAdmin bool
}

func NewAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewFailResponse("wrong token"))
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewFailResponse(err.Error()))
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewFailResponse(err.Error()))
			return
		}

		auth := CustomClaim{
			Id:      uint(claim["user_id"].(float64)),
			IsAdmin: claim["is_admin"].(bool),
		}
		c.Set("auth", auth)
	}
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token invalid")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GenerateToken(userId uint, isAdmin bool) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId
	claim["is_admin"] = isAdmin

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
