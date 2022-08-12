package keycloak_middleware

import "fmt"
import "github.com/labstack/echo/v4"

//EchoGuard set up the echo middleware and access
func (m keyCloakMiddleware) EchoGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			keyCloakENV := m.config

			accessToken := ctx.Request().Header.Get("Authorization")

			info, err := m.goCloak.RetrospectToken(ctx.Request().Context(), accessToken, keyCloakENV.ClientID, keyCloakENV.ClientSecret, keyCloakENV.Realm)
			if err != nil {
				return err
			}
			if *info.Active == false {
				return fmt.Errorf("invalid token")
			}

			_, claims, err := m.goCloak.DecodeAccessToken(ctx.Request().Context(), accessToken, keyCloakENV.Realm)
			if err != nil {
				return err
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
				return next(ctx)
			}

			return err
		}
	}
}
