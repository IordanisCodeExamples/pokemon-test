package data

import "errors"

type Data struct {
	DataInterface DataInterface
}

type DataInterface interface {
	SearcPokemon(name string) (string, error)
}

func NewData(dataInterface DataInterface) (*Data, error) {
	if dataInterface == nil {
		return nil, errors.New("data interface is required")
	}

	return &Data{
		DataInterface: dataInterface,
	},nil
}