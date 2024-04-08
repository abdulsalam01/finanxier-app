By Abdul Salam

## Introduction

This project is a RESTful API service written in Golang designed to manage products and user sessions. It ensures secure access to endpoints through JWT (JSON Web Tokens) and adheres to best practices in code conventions.

## Table of Contents

- [Tech Stack](#tech-stack)
- [Setup](#setup)
- [Manual Setup](#manual-setup)
- [Database Schema](#database-schema)
- [Endpoints](#endpoints)
- [Libraries Used](#libraries-used)
- [Configuration](#configuration)
- [Design Pattern](#design-pattern)
- [Documentation](#documentation)
- [Optimization Considerations](#optimization-considerations)

## Tech Stack

- Golang 1.21
- Redis
- PostgreSQL

## Setup

1. Ensure your local PostgreSQL is disabled (or kill it manually), for instance:
   ```bash
   sudo kill -9 PID
   ```
2. Ensure Docker is installed on your machine.
3. Run the following command:
   ```bash
   docker-compose up
   ```
4. You're all set!

## Manual Setup

1. Create a database called `finanxier-app` on PostgreSQL.
2. Set up PostgreSQL connection credentials to `postgres` as both the username and password.
3. Run `redis-server` on your local machine.
4. Execute the command: `make run-http`.
5. You're done!

## Database Schema

There are two tables:

1. **User**: `(id, username, password_hash)`, with indexing on username.
2. **Products**: `(id, name, price)`

## Endpoints

**Base URL:** `/api/v1`
**Base PORT:** `:8080`

1. **Get Product by ID:** `GET /product/{id}` with header `Authorization: Bearer <token>`
   Response:
     ```json
     {
         "data": {
             "id": "55506c8e-f046-4517-8c2e-b50b3c7a8f1a",
             "name": "Book",
             "price": 1000
         },
         "message": "Successfully executed",
         "success": true
     }
     ```

2. **Get Product by Params:** `GET /product` with header `Authorization: Bearer <token>` using query params as limit and offset.
   Response:
     ```json
     {
         "data": {
             "products": [
                 {
                     "id": "55506c8e-f046-4517-8c2e-b50b3c7a8f1a",
                     "name": "Book",
                     "price": 1000
                 },
                 {
                     "id": "2f8d3f11-2354-482e-8555-8c2c0c5dc459",
                     "name": "Laptop",
                     "price": 5000
                 }
             ],
             "total": 2
         },
         "message": "Successfully executed",
         "success": true
     }
     ```

3. **Create Product:** `POST /product` with header `Authorization: Bearer <token>`.
   - Body:
     - `name` -> should be more than or equal to 4 characters.
     - `price` -> number, greater than 0.
   - Response:
     ```json
     {
         "data": {
             "id": "55506c8e-f046-4517-8c2e-b50b3c7a8f1a",
             "name": "Book",
             "price": 1000
         },
         "message": "Successfully executed",
         "success": true
     }
     ```

4. **Get Current User:** `GET /user/current` with valid authorization token.
   Response:
     ```json
     {
         "data": {
             "id": "55506c8e-f046-4517-8c2e-b50b3c7a8f1a",
             "name": "admin"
         },
         "message": "Successfully executed",
         "success": true
     }
     ```

5. **(Bypass) Generate Valid Token:** `GET /token-generator` without any header.
    Response:
   ```json
   {
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTI2NTcxMTcsInVzZXIiOiJhZG1pbiJ9.l0E9WPd_IMXqE2yKoA9IhHc29aZObGDUjVrkxzSXI1g",
    "message": "Successfully executed",
    "success": true
   }
   ```

## Libraries Used

- `gomock` for unit testing.
- `chi` for lightweight routing.
- `uuid` to generate unique strings based on time.
- `pmx` as a driver for PostgreSQL.
- `redis` as a Redis wrapper.
- `redis locker` to prevent duplicate requests.
- `httprest` for rate limiting.
- `makefile` for base setup.

## Configuration

- Configuration is managed using a YAML file format to define all layers of the application (DB, Redis, app name, port, etc.).

## Design Pattern

- The application uses the repository pattern to handle and separate code into various layers, including handler, usecase, and repository levels, which then interact with the database.
- The handler sanitizes the input, the usecase executes the main logic and returns data, and the repository manages data interactions with the datastore.

## Documentation

Find all the documentation in the `/docs` directory. 

### Visual Flow Diagram

![Visual Flow Diagram](/docs/flow-pattern.drawio.png)

### Postman Collection

- Access the Postman collection [here](https://api.postman.com/collections/31649827-a3a60c63-311d-4a65-81cd-48d6a9234648?access_key=PMAT-01HTYHX06JXS0NMXFCR3KTNY28) to interact with the API's endpoints.


## Optimization Considerations

- Optimizations include implementing caching strategies to reduce database load, optimizing database queries to improve response times, and considering horizontal scaling to accommodate growth in user traffic and data volume.

By Abdul Salam