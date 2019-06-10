package main

import (
	"os"

	_ "github.com/lib/pq"
	httplib "gitlab.com/semestr-6/projekt-grupowy/backend/go-libs/http-lib"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/attributes"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/categories"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/configuration"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/energy_resources"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/files"
	"gitlab.com/semestr-6/projekt-grupowy/backend/obsluga-formularzy/units"
)

func main() {
	getEnv()
	routes := getRoutes()
	addEndpoints(routes)
	httplib.DeclareNewRouterWithAuthorizationAndStart(routes)
}

func getEnv() {
	configuration.ConnectionString = os.Getenv("CONN_STR")
}

func getRoutes() (routes httplib.Routes) {
	routes = append(routes, attributes.GetRoutes()...)
	routes = append(routes, categories.GetRoutes()...)
	routes = append(routes, energy_resources.GetRoutes()...)
	routes = append(routes, files.GetRoutes()...)
	//routes = append(routes, resources.GetRoutes()...)
	routes = append(routes, units.GetRoutes()...)
	return
}
