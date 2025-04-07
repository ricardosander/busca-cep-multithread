package main

import (
	"fmt"
	"os"
	"time"
)

func main()	{

	input, err := GetInput()
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	
	fmt.Println("Buscando o CEP", *input, "...")

	chViaCep := make(chan *Cep)
	chBrasilCep := make(chan *Cep)

	go ExecuteBrasilCep(chBrasilCep, *input)
	go ExecuteViaCep(chViaCep, *input)

	select {
	case bCep := <-chBrasilCep:
		ProcessResult(bCep)
	case vCep := <-chViaCep:
		ProcessResult(vCep)
	case <-time.After(1 * time.Second):
		fmt.Println("Tempo de espera excedido. Nenhum resultado encontrado.")
	}
}

func GetInput() (*string, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("cep não informado")
	}
	return &os.Args[1], nil
}

func ExecuteViaCep(ch chan *Cep, input string) {
	cep, err := BuscaViaCep(input)
	if err != nil {
		fmt.Println("Erro ao buscar CEP no ViaCEP:", err)
		return
	}
	ch <- cep
}

func ExecuteBrasilCep(ch chan *Cep, input string) {
	cep, err := BuscaBrasilCep(input)
	if err != nil {
		fmt.Println("Erro ao buscar CEP no Brasil API:", err)
		return
	}
	ch <- cep
}

func ProcessResult(cep *Cep) {
	if cep != nil {
		fmt.Println("\nResultado:")
		fmt.Println(*cep.Print())
	}
}
