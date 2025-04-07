package main

import "fmt"

type Cep struct {
	Street          string
	Neighborhood    string
	City 	  	    string
	State 	  	    string
	Cep             string
    Origin          string
}

func (c *Cep) Print() *string {
    result := fmt.Sprintf("%s\nBairro: %s\nCidade: %s\nEstado: %s\nCEP: %s\nOrigem: %s", c.Street, c.Neighborhood, c.City, c.State, c.Cep, c.Origin)
    return &result
}