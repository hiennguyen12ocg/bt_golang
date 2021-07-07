package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"bt/project/repository"

)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	listProducts := repository.GetAllProducts()
	  j, _ := json.Marshal(listProducts)
  

//   j, _ := json.MarshalIndent(listProducts, "", "  ")
  fmt.Fprint(w, string(j))
	// return 
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " Hello 0")
}

func Hello1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " Hello 1")
}

func HelloName(write http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	fmt.Fprintf(write, " Hello %s", name)
}
