package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tomiprasetyo/golang-bunkerservice-crud/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBunkerRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
