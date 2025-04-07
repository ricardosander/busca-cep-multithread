package main

import (
	"encoding/json"
)

type BrasilCep struct {
	Street	   	 string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City 	  	 string `json:"city"`
	State 	  	 string `json:"state"`
	Cep          string `json:"cep"`
}

func BuscaBrasilCep(cep string) (*Cep, error) {
	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	result, err := BuscaCEP(url)
	if err != nil {
		return nil, err
	}
	var brasilCep BrasilCep
	if err := json.Unmarshal([]byte(*result), &brasilCep); err != nil {
		return nil, err
	}
	return brasilCep.convertToCep(), nil
}

func (b *BrasilCep) convertToCep() *Cep {
    return &Cep{
        Street:       b.Street,
        Neighborhood: b.Neighborhood,
        City:        b.City,
        State:       b.State,
        Cep:         b.Cep,
        Origin:      "BrasilAPI",
    }
}
