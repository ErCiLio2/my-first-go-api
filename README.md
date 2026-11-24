# go-api

API REST em Go para gestão de produtos, construída com Gin e PostgreSQL, seguindo arquitetura em camadas (controller → usecase → repository).

## Tecnologias

- **Go 1.26**
- **Gin** — framework HTTP
- **PostgreSQL 15** — base de dados
- **Docker / Docker Compose**

## Estrutura

```
cmd/            # ponto de entrada (main.go)
controller/     # handlers HTTP
usecase/        # lógica de negócio
repository/     # acesso à base de dados
model/          # structs e tipos
db/             # ligação à base de dados
```

## Endpoints

| Método | Rota            | Descrição                  |
|--------|-----------------|----------------------------|
| GET    | /ping           | Verificação de disponibilidade |
| GET    | /products       | Listar todos os produtos   |
| POST   | /product        | Criar produto              |
| GET    | /product/:id    | Obter produto por ID       |
| PUT    | /product/:id    | Atualizar produto          |

## Executar com Docker

```bash
docker compose up --build
```

A API fica disponível em `http://localhost:8080`.

## Executar localmente

Requer PostgreSQL em execução. Ajusta as variáveis de ligação em `db/conn.go` e corre:

```bash
go run cmd/main.go
```
