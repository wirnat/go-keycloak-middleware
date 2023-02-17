package keycloak_middleware

import "github.com/Nerzal/gocloak/v12"

type keyCloakMiddleware struct {
	goCloak        gocloak.GoCloak
	config         KeyCloakConfig
	resourceAccess [][]string
	realmAccess    [][]string
}

func NewKeyCloakMiddleware(config KeyCloakConfig) *keyCloakMiddleware {
	goCloak := gocloak.NewClient(config.KeyCloakIP)
	return &keyCloakMiddleware{goCloak: *goCloak, config: config}
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
