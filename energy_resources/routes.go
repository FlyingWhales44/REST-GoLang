package energy_resources

import (
	"net/http"

	httplib "gitlab.com/semestr-6/projekt-grupowy/backend/go-libs/http-lib"
)

func GetRoutes() (routes httplib.Routes) {
	routes = httplib.Routes{
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-energy-resource-attribute",
			HandlerFunc: addEnergyResourceAttribute,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-energy-resource-attribute",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-gus-resource-attribute",
			HandlerFunc: addGUSResourceAttribute,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-gus-resource-attribute",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-all-energy-resource-attributes",
			HandlerFunc: getAllEnergyResourceAttributes,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-all-energy-resource-attributes",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-all-energy-resources",
			HandlerFunc: getAllEnergyResources,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-all-energy-resources",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-all-gus-resources",
			HandlerFunc: getAllGUSResources,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-all-gus-resources",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-gus-resources-id",
			HandlerFunc: getGUSResourcesID,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-gus-resources-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-energy-resource-attribute-by-id",
			HandlerFunc: getEnergyResourceAttributeById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-energy-resource-attribute-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-gus-resource-by-id",
			HandlerFunc: getGUSResourceById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-gus-resource-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-energy-resource-attribute",
			HandlerFunc: editEnergyResourceAttribute,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-energy-resource-attribute",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-gus-resource",
			HandlerFunc: editGUSResource,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-gus-resource",
			HandlerFunc: httplib.RespondStatus200,
		},
	}
	return
}
