package categories

import (
	"net/http"

	httplib "gitlab.com/semestr-6/projekt-grupowy/backend/go-libs/http-lib"
)

func GetRoutes() (routes httplib.Routes) {
	routes = httplib.Routes{
		httplib.Route{
			HttpMethod:  http.MethodPost,
			Route:       "/add-category",
			HandlerFunc: addCategory,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/add-category",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-categories",
			HandlerFunc: getCategories,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-categories",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-category-by-id",
			HandlerFunc: getCategoryById,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-category-by-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodGet,
			Route:       "/get-category-id",
			HandlerFunc: getCategoryId,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/get-category-id",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodPut,
			Route:       "/edit-category",
			HandlerFunc: editCategory,
		},
		httplib.Route{
			HttpMethod:  http.MethodOptions,
			Route:       "/edit-category",
			HandlerFunc: httplib.RespondStatus200,
		},
		httplib.Route{
			HttpMethod:  http.MethodDelete,
			Route:       "/delete-category",
			HandlerFunc: deleteCategory,
		},
	}
	return
}
