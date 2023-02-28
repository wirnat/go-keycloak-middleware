package keycloak_middleware

import (
	"context"
	"github.com/Nerzal/gocloak/v12"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Middleware interface {
	GetClaim(ctx context.Context, accessToken string) (claim *jwt.MapClaims, err error)
	ResourceAccess(args ...string) (r keyCloakMiddleware)
	RealmAccess(args ...string) (r keyCloakMiddleware)
	GinGuard(hook ...GinHook) gin.HandlerFunc
	EchoGuard(hook ...EchoHook) echo.MiddlewareFunc
	ValidateRealmAccess(claims jwt.MapClaims) (err error)
	ValidateResourceAccess(claims jwt.MapClaims) error
	ReturnGoCloak() *gocloak.GoCloak
}
