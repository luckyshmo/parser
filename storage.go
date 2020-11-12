package main

import (
	"errors"
)

type storage interface {
	insert(j Joint) error
	remove(number string) error
	get(number string) (Joint, error)
}

type jointStorage struct {
	data map[string]Joint
}

func newJointStorage() *jointStorage {
	return &jointStorage{
		data: make(map[string]Joint),
	}
}

func (js *jointStorage) insert(j Joint) error {
	js.data[j.number] = j

	return nil
}

func (js *jointStorage) get(number string) (Joint, error) {
	j, exists := js.data[number]
	if !exists {
		return Joint{}, errors.New("No joints with such number")
	}

	return j, nil
}

func (js *jointStorage) remove(number string) error {
	delete(js.data, number)
	return nil
}
