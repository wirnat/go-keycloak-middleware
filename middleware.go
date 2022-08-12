package keycloak_middleware

type Middleware interface {
	ResourceAccess(args ...string) (r Middleware)
	RealmAccess(args ...string) (r Middleware)
}
