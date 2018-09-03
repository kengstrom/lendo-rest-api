package main

import (
  "testing"
  "strconv"
  "net/http"
  "net/http/httptest"
)

var api API


// Deletes All Stores for test purposes
func DeleteAllStores() {
  stores = stores[:0]
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
      stores = append(stores, BikeStore{ID:strconv.Itoa(i), Name: "Store #"+strconv.Itoa(i) , Address: "Address "+strconv.Itoa(i) })
    }
}

func TestMain(m *testing.M) {
  api := API{}
  api.Run()
}

func TestGetStore(t *testing.T) {
  CreateSomeStores(1)
  request, _ := http.NewRequest("GET", "/bikestore/1", nil)
  response := executeRequest(request)
  checkResponseCode(t, http.StatusOK, response.Code)
}
