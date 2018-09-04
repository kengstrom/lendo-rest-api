package main

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

type API struct {
  Router *mux.Router
  stores []BikeStore
}


func (api *API) Run() {
  api.Router = mux.NewRouter()
  api.Router.HandleFunc("/bikestores", api.getBikeStores).Methods("GET")
  api.Router.HandleFunc("/bikestore/{id}", api.getBikeStore).Methods("GET")
  api.Router.HandleFunc("/bikestore", api.createBikeStore).Methods("POST")
  api.Router.HandleFunc("/bikestore/{id}", api.updateBikeStore).Methods("PUT")
  api.Router.HandleFunc("/bikestore/{id}", api.deleteBikeStore).Methods("DELETE")
}  

func (api *API) getBikeStores(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(stores)
}



func (api *API) getBikeStore(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  store := BikeStore{ ID: id}
  if err := store.getStore(); err != nil {
    response, _ := json.Marshal(map[string]string{"error": err.Error()})
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write(response)
  }
  response, _ := json.Marshal(store)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(response)
}


func (api *API) updateBikeStore(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var store BikeStore
    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()
    err := decoder.Decode(&store)
    if err != nil {
            
    }
    store.ID = vars["id"]
    store.updateStore()
    
}

func (api *API) createBikeStore(w http.ResponseWriter, r *http.Request) {
}


func (api *API) deleteBikeStore(w http.ResponseWriter, r *http.Request) {
}