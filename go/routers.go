/*
 * Inventory API
 *
 * This is a sample inventory management server based on the OpenAPI 3.0 specification
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"PostCategory",
		strings.ToUpper("Post"),
		"/category",
		PostCategory,
	},

	Route{
		"DeleteCategoryById",
		strings.ToUpper("Delete"),
		"/category/{id}",
		DeleteCategoryById,
	},

	Route{
		"GetCategories",
		strings.ToUpper("Get"),
		"/category",
		GetCategories,
	},

	Route{
		"GetCategoryById",
		strings.ToUpper("Get"),
		"/category/{id}",
		GetCategoryById,
	},

	Route{
		"PutCategoryById",
		strings.ToUpper("Put"),
		"/category/{id}",
		PutCategoryById,
	},

	Route{
		"PostContainer",
		strings.ToUpper("Post"),
		"/container",
		PostContainer,
	},

	Route{
		"DeleteContainerById",
		strings.ToUpper("Delete"),
		"/container/{id}",
		DeleteContainerById,
	},

	Route{
		"GetContainerById",
		strings.ToUpper("Get"),
		"/container/{id}",
		GetContainerById,
	},

	Route{
		"GetContainers",
		strings.ToUpper("Get"),
		"/container",
		GetContainers,
	},

	Route{
		"PutContainerById",
		strings.ToUpper("Put"),
		"/container/{id}",
		PutContainerById,
	},

	Route{
		"PostItem",
		strings.ToUpper("Post"),
		"/item",
		PostItem,
	},

	Route{
		"DeleteItemById",
		strings.ToUpper("Delete"),
		"/item/{id}",
		DeleteItemById,
	},

	Route{
		"GetItems",
		strings.ToUpper("Get"),
		"/item",
		GetItems,
	},

	Route{
		"GetItemById",
		strings.ToUpper("Get"),
		"/item/{id}",
		GetItemById,
	},

	Route{
		"PutItemById",
		strings.ToUpper("Put"),
		"/item/{id}",
		PutItemById,
	},
}
