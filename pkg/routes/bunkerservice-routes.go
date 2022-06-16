package routes

import (
	"github.com/gorilla/mux"
	"github.com/tomiprasetyo/golang-bunkerservice-crud/pkg/controllers"
)

// membuat routes untuk request
// anonymous function
var RegisterBunkerRoutes = func(r *mux.Router) {
	r.HandleFunc("/bunker/", controllers.CreateBunker).Methods("POST")
	r.HandleFunc("/bunker/", controllers.GetAllBunker).Methods("GET")
	r.HandleFunc("/bunker/{bunkerId}", controllers.GetBunkerById).Methods("GET")
	r.HandleFunc("/bunker/{bunkerId}", controllers.UpdateBunker).Methods("PUT")
	r.HandleFunc("/bunker/{bunkerId}", controllers.DeleteBunker).Methods("DELETE")
}
