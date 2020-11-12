package storage

import (
	"errors"
)

type storage interface {
	insert(j joint) error
	remove(number string) error
	get(number string) (joint, error)
}

type jointStorage struct {
	data map[string]joint
}

func newJointStorage() *jointStorage {
	return &jointStorage{
		data: make(map[string]joint),
	}
}

func (js *jointStorage) insert(j joint) error {
	js.data[j.number] = j

	return nil
}

func (js *jointStorage) get(number string) (joint, error) {
	j, exists := js.data[number]
	if !exists {
		return joint{}, errors.New("No joints with such number")
	}

	return j, nil
}

func (js *jointStorage) remove(number string) error {
	delete(js.data, number)
	return nil
}
