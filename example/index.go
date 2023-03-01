package main

import (
	"github.com/labstack/echo/v4"
	"github.com/wirnat/go-keycloak-middleware"
)

func main() {
	//TODO: INIT Middleware instance
	middleware := keycloak_middleware.NewKeyCloakMiddleware(keycloak_middleware.KeyCloakConfig{
		KeyCloakIP:   "localhost:8080",
		Realm:        "my-erp",
		ClientID:     "your client id",
		ClientSecret: "your client secret",
	})

	//TODO: Build a permission role
	adminAccess := middleware.RealmAccess("admin").ResourceAccess().EchoGuard()
	e := echo.New()
	//TODO: implement it to your route
	e.POST("login", func(c echo.Context) error {
		panic("something")
	}, adminAccess)
	e.Start(":8181")

	//TODO: DONE
}
