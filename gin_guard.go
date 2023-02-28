package keycloak_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wirnat/go-keycloak-middleware/response"
)

//GinGuard set up the gin middleware and access
func (m keyCloakMiddleware) GinGuard(ginHook ...GinHook) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyCloakENV := m.config

		accessToken := ctx.Request.Header.Get("Authorization")
		//info, err := m.goCloak.RetrospectToken(ctx.Request.Context(), accessToken, keyCloakENV.ClientID, keyCloakENV.ClientSecret, keyCloakENV.Realm)
		//if err != nil {
		//	response.UnauthorizedFailDetail("invalid token", err.Error(), ctx)
		//	return
		//}
		//if *info.Active == false {
		//	response.UnauthorizedFailDetail("invalid token", "your token has been inactive", ctx)
		//	return
		//}

		_, claims, err := m.GoCloak.DecodeAccessToken(ctx.Request.Context(), accessToken, keyCloakENV.Realm)
		if err != nil {
			response.UnauthorizedFailDetail("decode token failed", "your token is invalid", ctx)
			return
		}

		errs := []string{}

		for _, hook := range ginHook {
			err = hook(ctx, claims)
			if err != nil {
				response.UnauthorizedFailDetail(errs, err.Error(), ctx)
				return
			}
		}

		isValid := false

		if len(m.realmAccess) < 1 && len(m.resourceAccess) < 1 {
			isValid = true
		}

		err = m.ValidateResourceAccess(*claims)
		if err == nil {
			isValid = true
		} else {
			errs = append(errs, err.Error())
		}

		err = m.ValidateRealmAccess(*claims)
		if err == nil {
			isValid = true
		} else {
			errs = append(errs, err.Error())
		}

		if isValid {
			ctx.Next()
			return
		} else {
			response.UnauthorizedFailDetail(errs, "you doesn't have access to this endpoint", ctx)
			return
		}
	}
}
