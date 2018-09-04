package main_test

import (
  "testing"
  "strconv"
  "net/http"
  "net/http/httptest"
  "os"
  "encoding/json"
  "bytes"
  lendo "github/kengstrom/lendo-rest-api"
)

var api lendo.API


// Deletes All Stores for test purposes
func DeleteAllStores() {
  api.Provider.Stores = api.Provider.Stores[:0]
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    api.Router.ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

// Creates some stores for test purposes
func CreateSomeStores(nbr int) {
    for i := 0; i < nbr; i++ {
      api.Provider.Stores = append(api.Provider.Stores, lendo.BikeStore{ID:strconv.Itoa(i), Name: "Store #"+strconv.Itoa(i) , Address: "Address "+strconv.Itoa(i) })
    }
}

func TestMain(m *testing.M) {
  api.Setup()
  code := m.Run()
  os.Exit(code)
}

func TestGetStore(t *testing.T) {
  DeleteAllStores()
  CreateSomeStores(1)
  request, _ := http.NewRequest("GET", "/bikestore/0", nil)
  response := executeRequest(request)
  checkResponseCode(t, http.StatusOK, response.Code)

}

func TestCreateStore(t *testing.T) {
  DeleteAllStores()
  newStore := lendo.BikeStore{ID: "123", Name: "CreatedStore", Address: "CreatedAddress"}
  payload, _ := json.Marshal(newStore)
  req, _ := http.NewRequest("POST", "/bikestore", bytes.NewBuffer(payload))
	response := executeRequest(req)
  checkResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("GET", "/bikestore/3", nil)
	response = executeRequest(req)
	json.Unmarshal(response.Body.Bytes(), &newStore)
  if (newStore.Name != "CreatedStore") {
    t.Errorf("Expected store name %s. Got %s\n", "CreatedStore", newStore.Name)
  }

}

func TestUpdateStore(t *testing.T) {
  DeleteAllStores()
  CreateSomeStores(10)

  //first get a stores
  req, _ := http.NewRequest("GET", "/bikestore/3", nil)
	response := executeRequest(req)
	var orgStore lendo.BikeStore
	json.Unmarshal(response.Body.Bytes(), &orgStore)

  orgStore.Name = "NewBikeStoreName"

  payload, _ := json.Marshal(orgStore)
	req, _ = http.NewRequest("PUT", "/bikestore/3", bytes.NewBuffer(payload))
	response = executeRequest(req)
  checkResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("GET", "/bikestore/3", nil)
	response = executeRequest(req)
	var newStore lendo.BikeStore
	json.Unmarshal(response.Body.Bytes(), &newStore)
  if (newStore.Name != "NewBikeStoreName") {
    t.Errorf("Expected store name %s. Got %s\n", "NewBikeStoreName", newStore.Name)
  }
}

func TestDeleteStore(t *testing.T) {
  DeleteAllStores()
  CreateSomeStores(10)

  request, _ := http.NewRequest("DELETE", "/bikestore/5", nil)
  response := executeRequest(request)
  checkResponseCode(t, http.StatusOK, response.Code)

  request, _ = http.NewRequest("GET", "/bikestore/5", nil)
  response = executeRequest(request)
  checkResponseCode(t, http.StatusNotFound, response.Code)

}
