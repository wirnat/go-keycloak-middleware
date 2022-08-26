package keycloak_middleware

import "github.com/gin-gonic/gin"
import 	"github.com/golang-jwt/jwt/v4"

type GinHook func(ctx *gin.Context,claims *jwt.MapClaims )error
