package keycloak_middleware

type access struct {
	resourceAccess [][]string
	realmAccess    [][]string
}

//ResourceAccess validate resource access permission, args mean OR
func (m access) ResourceAccess(args ...string) (r access) {
	m.resourceAccess = append(m.resourceAccess, args)
	return m
}

//RealmAccess validate realm access permission, args mean OR
func (m access) RealmAccess(args ...string) (r access) {
	m.realmAccess = append(m.realmAccess, args)
	return m
}
