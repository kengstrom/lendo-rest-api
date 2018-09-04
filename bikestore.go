package main

import (
    "errors"
)


type BikeStore struct {
  ID    string    `json:"id,omitempty"`
  Name  string    `json:"name,omitempty"`
  Address string   `json:"address,omitempty"`
}


func (bs *BikeStore) getStore(dp* DataProvider) error {
  for _, item := range dp.Stores {
    if bs.ID == item.ID {
      *bs = item
      return nil
    }
  }
  return errors.New("Not found")
}

func (bs *BikeStore) createStore(dp* DataProvider) error {
  dp.Stores = append(dp.Stores, *bs)
  return nil
}

func (bs *BikeStore) updateStore(dp* DataProvider) error {
  for i, item := range dp.Stores {
    if bs.ID == item.ID {
      dp.Stores[i] = *bs
      return nil
    }
  }
  return errors.New("Not found")
}

func (bs *BikeStore) deleteStore(dp* DataProvider) error {
  for i, item := range dp.Stores {
    if bs.ID == item.ID {
      dp.Stores = append(dp.Stores[:i], dp.Stores[i+1:]...)
      return nil
    }
  }
  return errors.New("Not found")
}
