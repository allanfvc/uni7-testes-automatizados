# Testes automatizados
📜 Projeto para a disciplina de testes automatizados da especialização em práticas devops da uni7.

![Go](https://github.com/allanfvc/uni7-testes-automatizados/workflows/Go/badge.svg)

## 🛠 Instalação
* Primeiro baixe e instale o [Go](https://golang.org/dl/). `1.15` ou superior.
* Em seguida baixe as dependências:
    ```bash
    go get -v -t -d ./...
    ```
## 🚀 Executando os testes
### Testando um pacote específico
* No terminal, acesse o pacote a ser testado, por exemplo:
  ```bash
  cd calculator
  ```
* Em seguida execute o comando de teste:
  ```bash
  go test
  ```
  
### Testando o projeto inteiro
* Execute o comando de teste:
  ```bash
  go test -v ./...
  ```
