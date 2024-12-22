# Full Cycle - Terceiro Desafio - Listagem das orders

Este projeto implementa a listagem de orders utilizando: REST, GRPC e GraphQL.

## Funcionalidades

- Criação de orders
- Listagem de orders

## Pré-requisitos

**GO**

Certifique-se de ter o GO instalado em sua máquina, de preferência a versão 1.22 ou superior. Se você não tiver o GO instalado em sua máquina, você pode fazer isso, consultando a [documentação do GO](https://go.dev/doc/install) e seguindo os passos para a instalação do mesmo.


**Docker**

O projecto faz uso do Docker para trabalhar com banco de dados e os serviços do RabbitMQ então certifique-se de ter o Docker instalado em sua máquina, de preferência a versão mais atualizada. Se você não tiver o Docker instalado em sua máquina, você pode fazer isso, consultando a [documentação do Docker](https://docs.docker.com/desktop/) e seguindo os passos para a instalação do mesmo.

## Instalação

O projeto utiliza bibliotecas externas então, certifique-se de instalar as dependências. Acesse o diretório raiz no terminal e execute o comando abaixo:

```bash

git clone https://github.com/manueldembo/clean_architecture.git

cd clean_architecture

go mod tidy

```

### Docker-Compose

Para iniciar o banco de dados e os serviços do RabbitMQ acesse o diretório raiz `~clean_architecture`, acesse o terminal e execute o comando abaixo:


```bash
 docker-compose up -d
```

### Migrações

Para criar as migrações para o bando de dados acesse o diretório raiz `~clean_architecture`, acesse o terminal e execute o comando abaixo antes de executar a aplicação:


```bash
 migrate create -ext=sql -dir=sql/migrations -seq init

  migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" up
```

## Ambiente de desenvolvimento

Após concluir o processo de instalação, acesse terminal no diretório `cmd/ordersystem/` e execute o comando abaixo para iniciar os servidor.


```bash
 go run main.go wire_gen.go
```

### Nota!

Para garantir que aplicação será apresentada como esperado você precisa garantir que as variáveis de ambiente estão devidamente definidas. Abaixo as configurações para o ambiente de desenvolvimento:

```bash
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=:8000
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8080
```


## Endpoints

**REST**
- GET /order
- POST /orde

**GraphQL**
- Query `ListOrders`
- Mutation `CreateOrder`

**gRPC**
- Service `ListOrders`
- Service `CreateOrder`

## Portas por serviço

- **REST**      :8000
- **gRPC**      :50051
- **GraphQL**   :8080
