# Go API — Products Service

## Overview

Simple REST API for managing products (CRUD) using Gin and PostgreSQL.

- Endpoints:
  - `GET /ping` — health check
  - `GET /products` — list products
  - `POST /products` — create product
  - `GET /product/:productId` — get product by ID
  - `PATCH /product/:productId` — update product by ID (partial; note missing leading slash)
  - `DELETE /product/:productId` — delete product by ID (note missing leading slash)

## Prerequisites

- Docker and Docker Compose (recommended), or
- Go toolchain and a running PostgreSQL instance

## Configuration

Environment variables are loaded from `.env`:

```
DB_HOST=go_db
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_dbname
```

## Database Schema

Create the `product` table before running:

```sql
CREATE TABLE IF NOT EXISTS product (
  id SERIAL PRIMARY KEY,
  product_name TEXT NOT NULL,
  price NUMERIC(10,2) NOT NULL
);
```

## Run with Docker (recommended)

1. Build and start:
   ```
   docker-compose up --build
   ```
2. Once healthy, API will be at:
   - http://localhost:8080/ping
   - http://localhost:8080/products

## Run locally (without Docker)

1. Start PostgreSQL and ensure the table above exists.
2. Export environment (or keep `.env` in project root).
3. Install deps:
   ```
   go mod tidy
   ```
4. Run:
   ```
   go run cmd/main.go
   ```

## Using the API (examples)

For API examples, please refer to the official Postman documentation or the API documentation provided in the project.

- [Postman API Documentation](https://www.postman.com/altimetry-architect-54647672/workspace/rodrigo-s-public-apis/collection/25405760-dd97761c-2e74-4237-a95c-94e8e9d5425f?action=share&creator=25405760) (replace with your specific API link if available)

## VS Code

- Use the preconfigured launch settings in `.vscode/launch.json`.
- Choose "GO: Launch Package" to debug the whole app.

## Project Structure

- `cmd/main.go` — HTTP server wiring and route registration
- `controller/` — HTTP handlers
- `usecase/` — application business logic
- `repository/` — database access
- `model/` — data models
- `db/` — database connection
