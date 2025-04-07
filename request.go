package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func BuscaCEP(url string) (*string, error) {
	http.NewRequest("GET", url, nil)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao buscar o CEP: %s", resp.Status)
	}
	bResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sResponse := string(bResponse)
	return &sResponse, nil
}