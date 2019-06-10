package attributes

import (
	"net/http"

	httplib "gitlab.com/semestr-6/projekt-grupowy/backend/go-libs/http-lib"
)

func GetRoutes() (routes httplib.Routes) {
	routes = httplib.Routes{
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-factor-attribute",
			HandlerFunc: addFactorAttribute,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-factor-attribute",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-factor",
			HandlerFunc: addFactor,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-factor",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-attribute-to-existing-resource",
			HandlerFunc: addAttributeToExistingResource,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-attribute-to-existing-resource",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-resource",
			HandlerFunc: addResource,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-resource",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-all-factor-attributes",
			HandlerFunc: getAllFactorAttributes,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-all-factor-attributes",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-all-factors",
			HandlerFunc: getAllFactors,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-all-factors",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-factors-id",
			HandlerFunc: getFactorsId,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-factors-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-factor-by-id",
			HandlerFunc: getFactorById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-factor-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-resource-by-id",
			HandlerFunc: getResourceById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-resource-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-factor-attribute-by-id",
			HandlerFunc: getFactorAttributeById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-factor-attribute-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-resources-with-attributes",
			HandlerFunc: getResourcesWithAttributes,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-resources-with-attributes",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-factor-attribute",
			HandlerFunc: editFactorAttribute,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-factor-attribute",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-factor",
			HandlerFunc: editFactor,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-factor",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-resource",
			HandlerFunc: editResource,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-resource",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodDelete,
			Route:       "/delete-factor-resource",
			HandlerFunc: deleteFactorResource,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/delete-factor-resource",
			HandlerFunc: httplib.RespondStatus200,
		},
	}
	return
}
