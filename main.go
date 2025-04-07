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
		CloseChannels(chBrasilCep, chViaCep)
	case vCep := <-chViaCep:
		ProcessResult(vCep)
		CloseChannels(chBrasilCep, chViaCep)
	case <-time.After(1 * time.Second):
		fmt.Println("Tempo de espera excedido. Nenhum resultado encontrado.")
		CloseChannels(chBrasilCep, chViaCep)
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
		close(ch)
		return
	}
	ch <- cep
}

func ExecuteBrasilCep(ch chan *Cep, input string) {
	cep, err := BuscaBrasilCep(input)
	if err != nil {
		close(ch)
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

func CloseChannels(ch1, ch2 chan *Cep) {
	close(ch1)
	close(ch2)
}