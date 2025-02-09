# HiveLock
A lightweight secrets manager that securely stores and retrieves secrets, with an API for integration. It can be used for Kubernetes, CI/CD pipelines, or local development.

## Features

- Secure Storage - Encrypt secrets before storing them in a database (e.g., SQLite, PostgreSQL, or BoltDB).
- REST API - CRUD operations for secrets.
- Access Control - Role-based access or API tokens for authentication.
- Expiration & Rotation - Auto-expiring secrets and rotation policies.
- Kubernetes Integration - Sync secrets with Kubernetes Secrets.
- Audit Logging - Track who accessed which secret.
- CLI Tool - Interact with the secrets manager from the terminal.

## Tech Stack

- Golang - Core programming language
- Gin/Fiber - Web framework for REST API
- GORM - ORM for database interactions
- AES/GCM - Encryption for secret storage
- JWT/OAuth - Authentication
- Kubernetes API Client (client-go) - For syncing secrets


## Install Dependencies

Run the following command to install GORM and SQLite driver:
```bash
go get github.com/joho/godotenv
go get -u gorm.io/gorm gorm.io/driver/sqlite
go get github.com/gofiber/fiber/v2
```

Create a `.env` file in the root directory with the SECRET_KEY environment variables.

## How to use


Store a Secrets
```bash
curl -X POST "http://localhost:8080/secrets" \
     -H "Content-Type: application/json" \
     -d '{"key": "db_password", "value": "mySuperSecretPassword"}'
```
Retrieve a Secret
```bash
curl -X GET "http://localhost:8080/secrets/1"
```
Delete a Secret
```bash
curl -X DELETE "http://localhost:8080/secrets/1"
```
