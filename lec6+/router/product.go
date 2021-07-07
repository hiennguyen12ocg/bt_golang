package router


import (
	"net/http"

	"github.com/gorilla/mux"

	"bt/project/controller"
)
func ConfigRouterProduct() *mux.Router {

	routers := mux.NewRouter()
	routers.HandleFunc("/", controller.Hello).Methods(http.MethodGet)
	routers.Methods(http.MethodGet).Path("/hello").HandlerFunc(controller.Hello1)

	routers.Methods(http.MethodGet).Path("/products").HandlerFunc(controller.GetAllProducts)

	routers.Methods(http.MethodGet).Path("/hello/{name}").HandlerFunc(controller.HelloName)

	return routers
}