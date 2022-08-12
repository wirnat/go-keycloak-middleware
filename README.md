
## Acknowledgements

 - [Keycloak SSO](https://www.keycloak.org/)
 - [Go Cloak](https://github.com/Nerzal/gocloak)




## Installation

Install using go module in your project

```bash
  go get github.com/wirnat/go-keycloak-middleware
```

## Overview

- Simple middleware for validating role from keycloak claims

- Support for Echo handler

- Support for Gin handler

- Validate user using keycloak realm access

- Validate user using keycloak resource access
## Usage/Examples
    //1.INIT Middleware instance
	middleware := keycloak_middleware.NewKeyCloakMiddleware(keycloak_middleware.KeyCloakConfig{
		KeyCloakIP:   "localhost:8080",
		Realm:        "my-erp",
		ClientID:     "your client id",
		ClientSecret: "your client secret",
	})

	//2. Build a permission role
	adminAccess := middleware.RealmAccess("admin").ResourceAccess().EchoGuard()

	e := echo.New()

	//3. implement it to your route
	e.POST("login", func(c echo.Context) error {
		panic("something")
	}, adminAccess)

	e.Start(":8181")
