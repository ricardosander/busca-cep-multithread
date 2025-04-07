package main

import (
	"encoding/json"
)

type ViaCep struct {
	Logradouro	 string `json:"logradouro"`
	Bairro 	 string `json:"bairro"`
	Cidade	 string `json:"localidade"`
	Estado 	 string `json:"estado"`
	Uf 		 string `json:"uf"`
	Cep 	 string `json:"cep"`
}

func BuscaViaCep(cep string) (*Cep, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	result, err := BuscaCEP(url)
	if err != nil {
		return nil, err
	}
	var viaCep ViaCep
	if err := json.Unmarshal([]byte(*result), &viaCep); err != nil {
		return nil, err
	}
	return viaCep.convertToCep(), nil
}

func (v *ViaCep) convertToCep() *Cep {
    return &Cep{
        Street:       v.Logradouro,
        Neighborhood: v.Bairro,
        City:        v.Cidade,
        State:       v.Estado,
        Cep:         v.Cep,
        Origin:      "ViaCEP",
    }
}
