# Full Cycle - Terceiro Desafio - Listagem das orders

Este projeto implementa a listagem de orders utilizando: REST, GRPC e GraphQL.

## Funcionalidades

- Criação de orders
- Listagem de orders

## Pré-requisitos

**GO**

Certifique-se de ter o GO instalado em sua máquina, de preferência a versão 1.22 ou superior. Se você não tiver o GO instalado em sua máquina, você pode fazer isso, consultando a [documentação do GO](https://go.dev/doc/install) e seguindo os passos para a instalação do mesmo.


**Docker**

O projeto faz uso do Docker para trabalhar com banco de dados, serviços do RabbitMQ e para executar a própria aplicação  então certifique-se de ter o Docker instalado em sua máquina, de preferência a versão mais atualizada. Se você não tiver o Docker instalado em sua máquina, você pode fazer isso, consultando a [documentação do Docker](https://docs.docker.com/desktop/) e seguindo os passos para a instalação do mesmo.

### Executando a aplicação

Para iniciar a aplicação `~clean_architecture`, acesse o terminal e execute o comando abaixo:

```bash
 docker-compose up -d
```

### Nota!

Para garantir que aplicação será apresentada como esperado você precisa garantir que as variáveis de ambiente estão devidamente definidas. Abaixo as configurações para o ambiente de desenvolvimento:

```bash
DB_DRIVER=mysql
DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=orders
WEB_SERVER_PORT=8000
GRPC_SERVER_PORT=50051
GRAPHQL_SERVER_PORT=8080
RABBITMQ_HOST=rabbitmq
```


## Endpoints

**REST**
- GET /order
- POST /order

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
