package keycloak_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Middleware interface {
	ResourceAccess(args ...string) (r keyCloakMiddleware)
	RealmAccess(args ...string) (r keyCloakMiddleware)
	GinGuard(hook ...GinHook) gin.HandlerFunc
	EchoGuard(hook ...EchoHook) echo.MiddlewareFunc
	ValidateRealmAccess(claims jwt.MapClaims) (err error)
	ValidateResourceAccess(claims jwt.MapClaims) error
}
