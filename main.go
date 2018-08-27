package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type BikeStore struct {
  ID    string    `json:"id,omitempty"`
  Name  string    `json:"name,omitempty"`
  Address string   `json:"address,omitempty"`
}


var stores []BikeStore

func main() {
    GetPlaces()
    router := mux.NewRouter()
    router.HandleFunc("/bikestores", GetBikeStores).Methods("GET")
    router.HandleFunc("/bikestore/{id}", GetBikeStore).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}


func GetBikeStores(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(stores)
}

func GetBikeStore(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, item := range stores {
    if item.ID == params["id"] {
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
  json.NewEncoder(w).Encode(map[string]string{"error": "Invalid place id"})
}
