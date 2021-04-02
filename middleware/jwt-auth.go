package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthResp struct {
	Status  int
	Message string
	Result  interface{}
}

func AuthorizeJWT(jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := new(AuthResp)
		//get header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			resp.Message = "Unauthorized"
			resp.Status = http.StatusUnauthorized
			resp.Result = authHeader
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
		//validate token
		jwtString := strings.Split(authHeader, "Bearer ")[1]
		token, err := jwtService.ValidateToken(jwtString)
		if err != nil {
			resp.Message = "Unauthorized"
			resp.Status = http.StatusUnauthorized
			resp.Result = authHeader
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
		if token.Valid {
			_ = token.Claims.(jwt.MapClaims)
		} else {
			resp.Message = err.Error()
			resp.Status = http.StatusUnauthorized
			resp.Result = token
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

	}
}
