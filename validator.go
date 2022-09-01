package keycloak_middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func (m keyCloakMiddleware) ValidateResourceAccess(claims jwt.MapClaims) error {
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

func (m keyCloakMiddleware) ValidateRealmAccess(claims jwt.MapClaims) (err error) {
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
