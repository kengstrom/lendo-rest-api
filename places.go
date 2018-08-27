package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    _"time"
)

const key = "AIzaSyBT_pgfx2wewAxXAf49VxqlQ2XOTRKZxS0"
const location = "59.3324492,18.0641401"
const url = "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%s&radius=2000&type=bicycle_store&key=%s&%s"


type Response struct {
		Results          []PlacesSearchResult `json:"results,omitempty"`
		HTMLAttributions []string             `json:"html_attributions,omitempty"`
		NextPageToken    string               `json:"next_page_token,omitempty"`
}

type PlacesSearchResult struct {
  PlaceID   string `json:"place_id"`
	Name         string `json:"name"`
	Vicinity  string   `json:"vicinity"`
}


func GetPlaces() {
  pageToken := ""
  for {
    queryUrl := fmt.Sprintf(url, location, key, pageToken)
    response, err := http.Get(queryUrl)
    if err != nil {
          fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
          data, _ := ioutil.ReadAll(response.Body)
          respData := Response{}
          json.Unmarshal(data, &respData)

          for _, searchResult := range respData.Results {
            stores = append(stores, BikeStore{ID:searchResult.PlaceID, Name: searchResult.Name, Address:searchResult.Vicinity })
          }
          if respData.NextPageToken == "" {
            fmt.Println("No more data")
            break;
          }

          pageToken = "pagetoken=" + respData.NextPageToken

          //some sleeping since Google dont like me
          //time.Sleep(60 * time.Second)
    }
  }
}
