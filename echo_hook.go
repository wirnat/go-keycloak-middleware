package keycloak_middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type EchoHook func(ctx echo.Context, claims *jwt.MapClaims) error
