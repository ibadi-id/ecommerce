# Ecommerce Graphql Golang

Micro service solution for simple e-commerce system, dengan menggunakan <strong>golang</strong>, <strong>Graphql</strong> dan orm menggunakan <strong>Ent</strong>


## Table of Contents

* [Technology](#technology)
* [Installation](#installation)
* [Folder Structure](#folder)
* [Documentation](#documentation)
* [Routing](#routing)
* [Example Query](#example)
* [Test](#test)

## Technology

- **Clean Architecture**
- **Go Languange**: Build fast, reliable, and efficient software at scale. [more](https://go.dev/)
- **GraphQL API**: Get many resources in a single request and [more](https://graphql.org/)
- **ORM - ENT**: Simple, yet powerful ORM for modeling and querying data. [more](https://entgo.io/)
- **Route - Echo**: High performance, extensible, minimalist Go web framework[more](https://echo.labstack.com/)
- **Dockerize**: Optimized for deployments using Docker
- **Hot Reload - Air**: Live reload for Go apps. [more](https://github.com/cosmtrek/air)

## Installation

```
Dockerize

run:
$ docker compose build --no-cache
$ docker compose up


Local
prepare mysql db name `ecommerce`

run:
$ go mod tidy
$ make migrate_schema
$ make start

visit http://localhost:9090/playground
```

## Folder

### Struktur folder ecommerce

```
├── cmd
│  ├── app
│  └── migration
├── config
├── documentation     # Documentation File
├── ent
├── graph
├── migration         # Sql Migration File
├── pkg
│  ├── adapter
│  │ ├── controller   # Controller
│  │ ├── repository   # Specific implementaion of repository
│  │ └── resolver     # GraphQL resolvers
│  │ 
│  │ 
│  ├── entity
│  │ └── model        # Entity of model, (e.g. ent.Customer, ent.Product)
│  │ 
│  │ 
│  ├── infrastructure
│  │ ├── datastore    # MySQL configuration
│  │ ├── graphql      # GrahpQL configuration
│  │ └── router       # Echo router
│  │ 
│  │ 
│  ├── usecase
│  │ ├── repository   # Interface for adapter
│  │ └── usecase      # Usecase for application/business logic
```




## Documentation

### File documentation

```
Testing : ./documentation/cover.html
Architecture : ./documentation/acrhitecture.drawio
Postman : ./documentation/Ecommerce.postman_collection.json

```

## Routing

### Routing API

<details>
<summary>/query</summary>
- Endpoint Graphql
</details>
<details>
<summary>/playground</summary>
- Test Playground Graphql
</details>
</details>
<br>

## Example

### Example Query

```
query customers{
  customers{
    id
    name
    email
    phone
    createdAt
    updatedAt
  }
}

query product{
  product(id:1){
    name
    descriptions
    price
    sku
    stock
    createdAt
    updatedAt
  }
}

mutation createCustomer{
  createCustomer(input:{
    name:"tester",
    phone:"08511234567",
    email:"test@gmail.com"
  }){
    name
    id
    email
    phone
    createdAt
    updatedAt
  }
}
```

## Test

```
Set environment variable
APP_ENV = test

Check configuration test in ./config/config.test.yml

run:
$ go test ./pkg/adapter/repository/...

```