package keycloak_middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wirnat/go-keycloak-middleware/response"
	"strings"
)

//GinGuard set up the gin middleware and access
func (m keyCloakMiddleware) GinGuard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyCloakENV := m.config

		accessToken := ctx.Request.Header.Get("Title")

		info, err := m.goCloak.RetrospectToken(ctx.Request.Context(), accessToken, keyCloakENV.ClientID, keyCloakENV.ClientSecret, keyCloakENV.Realm)
		if err != nil {
			response.UnauthorizedFailDetail("invalid token", "your token has been inactive", ctx)
			return
		}
		if *info.Active == false {
			response.UnauthorizedFailDetail("invalid token", "your token has been inactive", ctx)
			return
		}

		_, claims, err := m.goCloak.DecodeAccessToken(ctx.Request.Context(), accessToken, keyCloakENV.Realm)
		if err != nil {
			response.UnauthorizedFailDetail("decode token failed", "your token is invalid", ctx)
			return
		}

		isValid := false
		err = m.validateResourceAccess(*claims)
		if err == nil {
			isValid = true
		}

		err = m.validateRealmAccess(*claims)
		if err == nil {
			isValid = true
		}

		if isValid {
			ctx.Next()
		}

		response.UnauthorizedFailDetail("forbidden access", "you doesn't have access to this endpoint", ctx)
		return
	}
}

func (m keyCloakMiddleware) validateResourceAccess(claims jwt.MapClaims) error {
	for _, clientRole := range m.resourceAccess {
		for _, s := range clientRole {
			data := strings.Split(s, ".")
			if len(data) < 2 {
				return fmt.Errorf("client not found, example: ms-item.write")
			}

			clientID := data[0]
			permittedRole := data[1]

			accessMap, ok := claims["resource_access"].(map[string]interface{})
			if !ok {
				continue
			}

			accessRole, ok := accessMap[clientID]
			if !ok {
				continue
			}

			listRole, _ := accessRole.(map[string]interface{})["roles"].([]interface{})
			for _, r := range listRole {
				role, _ := r.(string)
				if role == permittedRole {
					return nil
				}
			}
		}
	}
	return fmt.Errorf("forbidden access")

}

func (m keyCloakMiddleware) validateRealmAccess(claims jwt.MapClaims) (err error) {
	for _, clientRole := range m.realmAccess {
		for _, s := range clientRole {
			permittedRole := s

			accessMap, ok := claims["realm_access"].(map[string]interface{})
			if !ok {
				continue
			}

			listRole, _ := accessMap["roles"].([]interface{})
			for _, r := range listRole {
				role, _ := r.(string)
				if role == permittedRole {
					return nil
				}
			}
		}
	}

	return fmt.Errorf("forbidden access")
}
