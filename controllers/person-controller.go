package controllers

import (
	"encoding/json"
	
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hng/models"
	"github.com/hng/utils"
)

var NewPerson models.Person
func GetPersonById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    personId := vars["personId"]
    ID, err := strconv.ParseInt(personId, 0, 0)
    if err != nil {
        
        name := personId
        personDetails, _ := models.GetPersonByName(name)
        res, _ := json.Marshal(personDetails)
        w.Header().Set("Content-type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
    } else {
        personDetails, _ := models.GetPersonById(ID)
        res, _ := json.Marshal(personDetails)
        w.Header().Set("Content-type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
    }
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
    createPerson := &models.Person{}
    utils.ParseBody(r, createPerson)

    nameFromBody := createPerson.Name
    nameFromQuery := r.URL.Query().Get("name")

    if nameFromBody == "" && nameFromQuery == "" {
        http.Error(w, "Name parameter is missing", http.StatusBadRequest)
        return
    }

    if nameFromBody != "" {
        createPerson.Name = nameFromBody
    } else {
        createPerson.Name = nameFromQuery
    }

    b := createPerson.CreatePerson()
    res, _ := json.Marshal(b)
    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}


func DeletePerson(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    personId := vars["personId"]
    ID, err := strconv.ParseInt(personId, 0, 0)
    if err != nil {
        name := personId
        person := models.DeletePersonByName(name)
        res, _ := json.Marshal(person)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
    } else {
        person := models.DeletePerson(ID)
        res, _ := json.Marshal(person)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
    }
}


func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    var updatePerson = &models.Person{}
    utils.ParseBody(r, updatePerson)
    vars := mux.Vars(r)
    personId := vars["personId"]
    ID, err := strconv.ParseInt(personId, 0, 0)
    if err != nil {
        name := personId
        personDetails, _ := models.GetPersonByName(name)
        if updatePerson.Name != "" {
            personDetails.Name = updatePerson.Name
        }
        models.GetDB().Save(&personDetails) 
        res, _ := json.Marshal(personDetails)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
    } else {
        personDetails, _ := models.GetPersonById(ID)
        if updatePerson.Name != "" {
            personDetails.Name = updatePerson.Name
        }
        models.GetDB().Save(&personDetails) 
        res, _ := json.Marshal(personDetails)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
    }
}


