package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hyusuri/golang_api/helper"
	"github.com/hyusuri/golang_api/service"
	"log"
	"net/http"
)

func AuthorizeJWT(jwtServ service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthHeader := c.GetHeader("Authorization")
		if AuthHeader == "" {
			response := helper.BuildErrorResponse("Failed to process", "No token", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtServ.ValidateToken(AuthHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id] :", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)

		}
	}
}
