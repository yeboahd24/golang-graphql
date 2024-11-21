# Enterprise Authentication System Setup Guide

This guide will walk you through setting up the Enterprise Authentication System from scratch.

## Prerequisites

1. Go 1.22 or later
2. PostgreSQL 12 or later
3. Git

## Project Setup

1. Create a new directory and initialize Go module:
```bash
mkdir enterprise-auth
cd enterprise-auth
go mod init auth-system
```

2. Install required dependencies:
```bash
# Web framework and GraphQL
go get github.com/gin-gonic/gin
go get github.com/99designs/gqlgen
go get github.com/vektah/gqlgen/cmd/gqlgen

# Database
go get gorm.io/gorm
go get gorm.io/driver/postgres

# Authentication
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt

# Configuration
go get github.com/joho/godotenv
```

3. Create project structure:
```bash
mkdir -p cmd/server
mkdir -p internal/{auth,database,middleware,models}
mkdir -p graph/{generated,model}
mkdir configs
```

4. Create `.env` file in the root directory:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=auth_system
JWT_SECRET=your-secret-key
PORT=8080
```

5. Set up PostgreSQL database:
```sql
CREATE DATABASE auth_system;
```

## GraphQL Setup

1. Create GraphQL schema in `graph/schema.graphql`:
```graphql
type User {
  id: ID!
  email: String!
  firstName: String
  lastName: String
  createdAt: String!
  updatedAt: String!
}

type AuthResponse {
  token: String!
  user: User!
}

input RegisterInput {
  email: String!
  password: String!
  firstName: String!
  lastName: String!
}

input LoginInput {
  email: String!
  password: String!
}

type Query {
  me: User!
}

type Mutation {
  register(input: RegisterInput!): AuthResponse!
  login(input: LoginInput!): AuthResponse!
}
```

2. Configure gqlgen by creating `gqlgen.yml`:
```yaml
schema:
  - graph/*.graphql

exec:
  filename: graph/generated/generated.go
  package: generated

model:
  filename: graph/model/models_gen.go
  package: model

resolver:
  layout: follow-schema
  dir: graph
  package: graph
  filename_template: "{name}.resolvers.go"

autobind:
  - "auth-system/graph/model"

models:
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.ID
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Int:
    model:
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
```

3. Generate GraphQL code:
```bash
go run github.com/99designs/gqlgen generate
```

## Implementation Steps

1. Create database models in `internal/models/user.go`
2. Set up database connection in `internal/database/postgres.go`
3. Implement authentication service in `internal/auth/service.go`
4. Create authentication middleware in `internal/middleware/auth.go`
5. Implement GraphQL resolvers in `graph/schema.resolvers.go`
6. Set up main server in `cmd/server/main.go`

## Running the Server

1. Make sure PostgreSQL is running and accessible
2. Update `.env` with your database credentials
3. Run the server:
```bash
go run cmd/server/main.go
```

The server will start at http://localhost:8080 (or the port specified in `.env`).
GraphQL playground will be available at http://localhost:8080

## Development Tools

- VS Code with Go extension is recommended
- GraphQL Playground for testing queries
- pgAdmin or similar tool for database management

## Security Considerations

1. Always use HTTPS in production
2. Store sensitive data (JWT_SECRET, DB_PASSWORD) securely
3. Implement rate limiting for API endpoints
4. Regular security audits
5. Keep dependencies updated
