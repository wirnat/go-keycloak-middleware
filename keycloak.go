package keycloak_middleware

import (
	"context"
	"github.com/Nerzal/gocloak/v12"
	"github.com/golang-jwt/jwt/v4"
)

type keyCloakMiddleware struct {
	GoCloak        *gocloak.GoCloak
	config         KeyCloakConfig
	resourceAccess [][]string
	realmAccess    [][]string
}

func NewKeyCloakMiddleware(config KeyCloakConfig, options ...func(cloak *gocloak.GoCloak)) *keyCloakMiddleware {
	goCloak := gocloak.NewClient(config.KeyCloakIP, options...)
	return &keyCloakMiddleware{GoCloak: goCloak, config: config}
}

func (m keyCloakMiddleware) ReturnGoCloak() *gocloak.GoCloak {
	return m.GoCloak
}

//ResourceAccess validate resource access permission, args mean OR
func (m keyCloakMiddleware) ResourceAccess(args ...string) (r keyCloakMiddleware) {
	m.resourceAccess = append(m.resourceAccess, args)
	return m
}

//RealmAccess validate realm access permission, args mean OR
func (m keyCloakMiddleware) RealmAccess(args ...string) (r keyCloakMiddleware) {
	m.realmAccess = append(m.realmAccess, args)
	return m
}

//GetClaim return claim from token
func (m keyCloakMiddleware) GetClaim(ctx context.Context, accessToken string) (claim *jwt.MapClaims, err error) {
	_, claims, err := m.GoCloak.DecodeAccessToken(ctx, accessToken, m.config.Realm)
	if err != nil {
		return nil, err
	}

	return claims, nil

}
