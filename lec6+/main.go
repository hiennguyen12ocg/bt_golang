package main

import (

	"net/http"

	"bt/project/connect"
	
	"bt/project/router"

	// "bt/project/repository"
)

func main() {
	routers := router.ConfigRouterProduct()
	
	//test connect database
	connect.Connect()
	


	//insert data
	// repository.InitData()

	//test get products
	// repository.GetAllProducts()


	err := http.ListenAndServe(":3000", routers)
	// log.Fatal(http.ListenAndServe(":3000", routers))
	if err != nil {
		panic(err)
	}
}




