package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTAuthMiddleware(secret string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SuccessHandler: nil,
		SigningKey:     []byte(secret),
		Claims:         &jwt.StandardClaims{},
	})
}
