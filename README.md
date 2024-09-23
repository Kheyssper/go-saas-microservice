# SAAS Microserice Platform Documentation

## Overview

This microservice manages platforms for a Software as a Service (SaaS) application. It provides endpoints to create, list, and manage platforms, with access control for platform creators and administrators.

## Environment Variables

Before running the service, set the following environment variables in a `.env` file:

```
POSTGRES_USER=your_username       # Database username
POSTGRES_PASSWORD=your_password     # Database password
POSTGRES_DB=saas_db                 # Database name
POSTGRES_HOST=localhost              # Database host
POSTGRES_PORT=5432                   # Database port
SERVER_PORT=8080                     # Port for the HTTP server
```

## Endpoints

- **List All Platforms**
  - **GET** `/platforms`
  - Retrieves a list of all platforms.

- **Create Platform**
  - **POST** `/platforms`
  - Creates a new platform. The request body should include:
    ```json
    {
      "platform_name": "example_name",
      "platform_slug": "example_slug",
    "creator_id":id
    }
    ```

- **Get Platform by ID**
  - **GET** `/platforms/{id}`
  - Retrieves a specific platform by its ID.

- **Run/Stop Platform**
  - **PUT** `/platforms/{id}/status`
  - Updates the status of a specific platform.

- **Delete Platform**
  - **DELETE** `/platforms/{id}`
  - Deletes a specific platform.

## Database Migration Command

To set up your database schema, run the following command, replacing placeholders with your actual PostgreSQL username and password:

```bash
psql -U your_username -d saas_db -f migrations/001_create_platforms_table.sql
```

## Running the Service

To run the microservice, execute:

```bash
go run cmd/platform_service/main.go
```

Ensure you have all environment variables set before starting the service.
