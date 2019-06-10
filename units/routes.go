package units

import (
	"net/http"

	httplib "gitlab.com/semestr-6/projekt-grupowy/backend/go-libs/http-lib"
)

func GetRoutes() (routes httplib.Routes) {
	routes = httplib.Routes{
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-unit",
			HandlerFunc: addUnit,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-unit",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-basic-unit",
			HandlerFunc: addBasicUnit,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-basic-unit",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-quantity",
			HandlerFunc: addQuantity,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-quantity",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-all-basic-units",
			HandlerFunc: getAllBasicUnits,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-all-basic-units",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-all-quantities",
			HandlerFunc: getAllQuantities,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-all-quantities",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-all-units",
			HandlerFunc: getAllUnits,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-all-units",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-basic-units-id",
			HandlerFunc: getBasicUnitsId,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-basic-units-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-quantity-id",
			HandlerFunc: getQuantityId,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-quantity-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-units-id",
			HandlerFunc: getUnitsId,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-units-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-basic-unit-by-id",
			HandlerFunc: getBasicUnitById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-basic-unit-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-unit-by-id",
			HandlerFunc: getUnitById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-unit-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-quanity-by-id",
			HandlerFunc: getQuanityById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-quanity-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-basic-unit",
			HandlerFunc: editBasicUnit,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-basic-unit",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-quantity",
			HandlerFunc: editQuantity,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-quantity",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-unit",
			HandlerFunc: editUnit,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-unit",
			HandlerFunc: httplib.RespondStatus200,
		},
	}
	return
}
