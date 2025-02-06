# Lab Golang with Mongo and Clean Arch

Este projeto é um laboratório para explorar a integração do Golang com o MongoDB utilizando a arquitetura limpa.

## Estrutura do Projeto

```
 - api/ 
     - create_order.http 
     - list_order.http 
 - cmd/ 
    - server/ 
        - main.go 
  - configs/ 
    - config.go 
   - internal/ 
        - entity/ 
             - interface.go 
             - order.go 
        - infra/ 
             - database/ 
                - mongo/ 
                 - connection/ 
                  - repository/ 
                   - schemas/ 
            - di/ 
                 - order_usecase.go 
             - enum/ 
                 - collections.go 
             - web/ 
                 - handler/ 
                     - handler_config.go
                      - order_handler.go 
                     - server/ 
       - usecase/ 
          - order_usecase.go  
          
- pkg/ 
    - errors/ 
        - main.go
``` 


## Configuração

### Arquivo de Configuração

O arquivo de configuração está localizado em [configs/config.go](configs/config.go). Ele utiliza a biblioteca `viper` para carregar as configurações do arquivo `.env`.

### Dependências

As dependências do projeto estão listadas no arquivo [go.mod](go.mod).

## Executando o Projeto

Para executar o projeto, siga os passos abaixo:

1. Certifique-se de ter o MongoDB rodando e configure as variáveis de ambiente no arquivo `.env`.
2. Execute o comando abaixo para iniciar o servidor:

```sh
go run cmd/server/main.go
``` 

### Endpoints da API

#### Criar Pedido

 - URL: /orders
 - Método: POST
 - Exemplo de Requisição:

```
POST http://localhost:8080/orders HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "price": 2733,
    "tax": 81
}
``` 

#### Listar Pedidos

 - URL: /orders
 - Método: GET
 - Exemplo de Requisição:

 ````
 GET http://localhost:8080/orders HTTP/1.1
Host: localhost:8000
````

### Estrutura de Diretórios
 - cmd/server/main.go: Ponto de entrada da aplicação.
 - configs/config.go: Carregamento das configurações.
 - internal/entity: Definições das entidades e interfaces.
 - internal/infra/database/mongo: Conexão e repositórios do MongoDB.
 - internal/infra/web: Handlers e servidor web.
 - internal/usecase: Casos de uso da aplicação.
 - pkg/errors: Tratamento de erros.