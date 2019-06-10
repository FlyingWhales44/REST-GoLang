package files

import (
	"net/http"

	httplib "gitlab.com/semestr-6/projekt-grupowy/backend/go-libs/http-lib"
)

func GetRoutes() (routes httplib.Routes) {
	routes = httplib.Routes{
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-source",
			HandlerFunc: addSource,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-source",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-source",
			HandlerFunc: getSource,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-source",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-source-by-id",
			HandlerFunc: getSourceById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-source-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-all-sources",
			HandlerFunc: getAllSource,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-all-sources",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-source",
			HandlerFunc: editSource,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-source",
			HandlerFunc: httplib.RespondStatus200,
		},
	}
	return
}
