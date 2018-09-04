package main

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  "log"
)

type API struct {
  Router *mux.Router
  //stores []BikeStore
  Provider *DataProvider
}

func (api *API) Setup() {
  api.Provider = &DataProvider{}

  api.Router = mux.NewRouter()
  api.Router.HandleFunc("/bikestores", api.getBikeStores).Methods("GET")
  api.Router.HandleFunc("/bikestore/{id}", api.getBikeStore).Methods("GET")
  api.Router.HandleFunc("/bikestore", api.createBikeStore).Methods("POST")
  api.Router.HandleFunc("/bikestore/{id}", api.updateBikeStore).Methods("PUT")
  api.Router.HandleFunc("/bikestore/{id}", api.deleteBikeStore).Methods("DELETE")
}
func (api *API) Run() {
  go api.Provider.GetPlaces()
  log.Fatal(http.ListenAndServe(":8000", api.Router))
}

func (api *API) getBikeStores(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(api.Provider.Stores)
}



func (api *API) getBikeStore(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  store := BikeStore{ ID: id}
  if err := store.getStore(api.Provider); err != nil {
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
      response, _ := json.Marshal(map[string]string{"error": err.Error()})
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusBadRequest)
      w.Write(response)
    }
    store.ID = vars["id"]
    if err := store.updateStore(api.Provider); err != nil {
      response, _ := json.Marshal(map[string]string{"error": err.Error()})
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusInternalServerError)
      w.Write(response)
    }

    response, _ := json.Marshal(store)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)

}

func (api *API) createBikeStore(w http.ResponseWriter, r *http.Request) {
  var store BikeStore
  decoder := json.NewDecoder(r.Body)
  defer r.Body.Close()
  err := decoder.Decode(&store)
  if err != nil {
    response, _ := json.Marshal(map[string]string{"error": err.Error()})
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusBadRequest)
    w.Write(response)
  }
  if err := store.createStore(api.Provider); err != nil {
    response, _ := json.Marshal(map[string]string{"error": err.Error()})
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusInternalServerError)
    w.Write(response)
  }
  response, _ := json.Marshal(store)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(response)

}


func (api *API) deleteBikeStore(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  store := BikeStore{ ID: id}
  if err := store.deleteStore(api.Provider); err != nil {
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
