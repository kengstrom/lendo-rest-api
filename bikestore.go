package main

import (
    "errors"
)


type BikeStore struct {
  ID    string    `json:"id,omitempty"`
  Name  string    `json:"name,omitempty"`
  Address string   `json:"address,omitempty"`
}


func (bs *BikeStore) getStore() error {
  for _, item := range stores {
    if bs.ID == item.ID {
      *bs = item
      return nil
    }
  }
  return errors.New("Not found")
}

func (bs *BikeStore) createStore() error {
  return errors.New("Not implemented")
}

func (bs *BikeStore) updateStore() error {
  return errors.New("Not implemented")
}

func (bs *BikeStore) deleteStore() error {
  return errors.New("Not implemented")
}
