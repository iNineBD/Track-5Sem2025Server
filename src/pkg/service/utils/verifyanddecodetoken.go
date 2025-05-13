package utils

import (
	"fmt"
	"inine-track/pkg/middleware"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyAndDecodeToken(c *gin.Context) (claims jwt.MapClaims, err error) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, fmt.Errorf("eror ao capturar o jwt token")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err = middleware.DecoteTokenJWT(token)

	if err != nil {
		return nil, err
	}

	return claims, nil
}
