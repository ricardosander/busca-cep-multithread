# Busca CEP Multithread
Projeto feito como desafio do módulo de Multithread da Pós Graduação Go Expert da Faculdade Full Cycle.

## Funcionamento
Ao rodar o programa passando um CEP como input, o mesmo irá fazer uma busca simultânea nas APIs [ViaCEP](https://brasilapi.com.br/) e [CepBrasil](https://viacep.com.br/), retornando o resultado da primeira a responder. Caso nenhuma das APIs responda em até 1s, o programa é abortado.

## Executando

Podemos rodar o programa no modo desenvolvimento:
```
go run ./... cep-informado
```

Ou buildar o programa e executá-lo:
```
go build
./busca-cep-multithread cep-informado
```