# Enterprise Authentication System API Documentation

## Overview

This API provides authentication and user management functionality through GraphQL. All endpoints are accessible through a single GraphQL endpoint at `/query`.

## Base URL

```
http://localhost:8080
```

GraphQL Playground: `http://localhost:8080`

## Authentication

Authentication is handled using JWT (JSON Web Tokens). After successful login or registration, you'll receive a token that should be included in subsequent requests.

### Headers

For authenticated requests, include the token in the Authorization header:

```
Authorization: Bearer <your_token>
```

## API Operations

### Mutations

#### Register User

Creates a new user account.

```graphql
mutation {
  register(input: {
    email: "user@example.com"
    password: "securepassword"
    firstName: "John"
    lastName: "Doe"
  }) {
    token
    user {
      id
      email
      firstName
      lastName
      createdAt
      updatedAt
    }
  }
}
```

Response:
```json
{
  "data": {
    "register": {
      "token": "eyJhbGciOiJIUzI1NiIs...",
      "user": {
        "id": "1",
        "email": "user@example.com",
        "firstName": "John",
        "lastName": "Doe",
        "createdAt": "2024-02-20T10:00:00Z",
        "updatedAt": "2024-02-20T10:00:00Z"
      }
    }
  }
}
```

#### Login

Authenticates a user and returns a JWT token.

```graphql
mutation {
  login(input: {
    email: "user@example.com"
    password: "securepassword"
  }) {
    token
    user {
      id
      email
      firstName
      lastName
      createdAt
      updatedAt
    }
  }
}
```

Response:
```json
{
  "data": {
    "login": {
      "token": "eyJhbGciOiJIUzI1NiIs...",
      "user": {
        "id": "1",
        "email": "user@example.com",
        "firstName": "John",
        "lastName": "Doe",
        "createdAt": "2024-02-20T10:00:00Z",
        "updatedAt": "2024-02-20T10:00:00Z"
      }
    }
  }
}
```

### Queries

#### Get Current User (Me)

Retrieves the profile of the currently authenticated user. Requires authentication.

```graphql
query {
  me {
    id
    email
    firstName
    lastName
    createdAt
    updatedAt
  }
}
```

Response:
```json
{
  "data": {
    "me": {
      "id": "1",
      "email": "user@example.com",
      "firstName": "John",
      "lastName": "Doe",
      "createdAt": "2024-02-20T10:00:00Z",
      "updatedAt": "2024-02-20T10:00:00Z"
    }
  }
}
```

## Error Handling

The API returns GraphQL errors in the following format:

```json
{
  "errors": [
    {
      "message": "Error message here",
      "path": ["fieldName"],
      "extensions": {
        "code": "ERROR_CODE"
      }
    }
  ]
}
```

Common error codes:
- `UNAUTHORIZED`: Authentication required or token invalid
- `INVALID_CREDENTIALS`: Invalid email or password
- `USER_EXISTS`: Email already registered
- `VALIDATION_ERROR`: Invalid input data
- `INTERNAL_ERROR`: Server error

## Rate Limiting

The API implements rate limiting to protect against abuse:
- Registration: 3 requests per IP per hour
- Login: 5 requests per IP per minute
- Other endpoints: 100 requests per IP per minute

## Security Notes

1. Always use HTTPS in production
2. Store tokens securely
3. Tokens expire after 24 hours
4. Implement proper error handling
5. Validate all inputs
6. Use strong passwords

## Testing in GraphQL Playground

1. Open `http://localhost:8080`
2. Use the documentation explorer (Docs) on the right
3. For authenticated requests, click "HTTP HEADERS" at the bottom and add:
```json
{
  "Authorization": "Bearer your_token_here"
}
```

## Best Practices

1. Always handle errors gracefully
2. Implement proper logging
3. Use HTTPS in production
4. Keep dependencies updated
5. Regular security audits
6. Monitor API usage
