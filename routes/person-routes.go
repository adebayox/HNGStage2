package routes

import (
	"github.com/gorilla/mux"
	"github.com/hng/controllers"
)

var RegisterPersonRoutes =func(router *mux.Router){
	router.HandleFunc("/api", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/api/{personId}", controllers.GetPersonById).Methods("GET")
	router.HandleFunc("/api/{personId}", controllers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/api/{personId}", controllers.DeletePerson).Methods("DELETE")
}