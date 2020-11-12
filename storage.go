package main

import (
	"errors"
	"parser/welding"
)

type storage interface {
	insert(j welding.Joint) error
	remove(number string) error
	get(number string) (welding.Joint, error)
}

type jointStorage struct {
	data map[string]welding.Joint
}

func newJointStorage() *jointStorage {
	return &jointStorage{
		data: make(map[string]welding.Joint),
	}
}

func (js *jointStorage) insert(j welding.Joint) error {
	js.data[j.Number] = j

	return nil
}

func (js *jointStorage) get(number string) (welding.Joint, error) {
	j, exists := js.data[number]
	if !exists {
		return welding.Joint{}, errors.New("No joints with such number")
	}

	return j, nil
}

func (js *jointStorage) remove(number string) error {
	delete(js.data, number)
	return nil
}
