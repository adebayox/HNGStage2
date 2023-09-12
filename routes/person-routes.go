package routes

import (
	"github.com/gorilla/mux"
	"github.com/hng/controllers"
)

var RegisterPersonRoutes =func(router *mux.Router){
	router.HandleFunc("/user", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/user/{personId}", controllers.GetPersonById).Methods("GET")
	router.HandleFunc("/user/{personId}", controllers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/user/{personId}", controllers.DeletePerson).Methods("DELETE")
}