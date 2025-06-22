# MBA LAB 2 - CEP & Weather Service

Este projeto consiste em dois microsserviços escritos em Go: um serviço de consulta de CEP (`cep-service`) e um serviço de consulta de clima (`weather-service`). Ambos utilizam instrumentação com OpenTelemetry e podem ser executados via Docker Compose.

## Estrutura do Projeto

```
cep-service/
    api/
    cmd/
    internal/
pkg/
otel/
weather-service/
    api/
    cmd/
    internal/
.gitignore
docker-compose.yaml
go.mod
README.md
```


## Serviços

### 1. cep-service

- Consulta informações de endereço a partir de um CEP usando a API ViaCEP.
- Endpoint principal: `POST /`  
  Exemplo de requisição:
  ```json
  {
    "cep": "29902555"
  }

### 2. weather-service

Consulta informações de clima a partir de uma localidade.
Endpoint principal: GET /{cep}
Exemplo de requisição:

```json
GET http://localhost:8081/18050605
```


## Como Executar
### Pré-requisitos

 - Docker
 - Docker Compose

#### Subindo os serviços

1. Execute:

```bash
docker-compose up --build
```

2. Acesse os endpoints:

 - CEP: http://localhost:8080/
 - Weather: http://localhost:8081/{cep}

#### Observabilidade
 - O projeto utiliza o Zipkin para rastreamento distribuído.
 - Acesse o Zipkin em: http://localhost:9411

#### Testes
Os testes unitários podem ser executados com:

```go
go test ./...
```