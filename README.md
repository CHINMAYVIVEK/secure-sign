
# Secure Sign

Secure Sign is a Golang-based authentication service that facilitates user registration, login, and retrieval of user details. The project includes REST API endpoints for seamless integration into your applications.

## Table of Contents

- [Secure Sign](#secure-sign)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Usage](#usage)
    - [API Endpoints](#api-endpoints)
    - [Docker Deployment](#docker-deployment)
  - [Project Structure](#project-structure)
    - [Profiling :](#profiling-)

## Getting Started

### Prerequisites

Before you start, ensure you have the following installed on your system:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Postgresql](https://www.postgresql.org/)
  
### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/CHINMAYVIVEK/secure-sign.git
   ```

2. Change into the project directory:

   ```bash
   cd secure-sign
   ```

3. Initialize the Go module:

   ```bash
   go mod init secure-sign
   ```

4. Build the Docker image:

   ```bash
   docker compose up
   ```

## Usage

### API Endpoints

Secure Sign provides the following API endpoints:

- **Register User:**

  `POST /register`

  Example Request:
  ```bash
  curl -X POST http://localhost:8080/register -d '{
    "first_name": "John",
    "last_name": "Doe",
    "gender": "Male",
    "age": 30,
    "email": "john.doe2@example.com",
    "phone_number": "9234567892",
    "salary": 50000,
    "password": "1234"}'
  ```

- **Login:**

  `POST /login`

  Example Request:
  ```bash
  curl -X POST http://localhost:8080/login -d '{"username": "newuser", "password": "password123"}'
  ```

- **Get User Details:**

  `GET /user/{username}`

  Example Request:
  ```bash
  curl http://localhost:8080/user/newuser
  ```

### Docker Deployment

To deploy the Secure Sign authentication service using Docker:

1. Build the Docker image as mentioned in the [Installation](#installation) section.

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 secure-sign
   ```

   The service will be accessible at [http://localhost:8080](http://localhost:8080).


## Project Structure

```

├── app
│   ├── health
│   │   └── health.go
│   └── user
│       ├── fetch_user.go
│       ├── login.go
│       ├── registration.go
│       └── user.go
├── config
│   ├── config.go
│   └── postgres.go
├── data
│   ├── data.sql
│   ├── secure-login.postman_collection.json
│   └── test.sql
├── helper
│   ├── helper.go
│   ├── postgres.go
│   ├── validation.go
│   └── zap-logger.go
├── logs
│   └── access.log
├── middleware
│   └── middleware.go
├── server
│   ├── router.go
│   └── server.go
├── test
│   ├── cpu.prof
│   ├── env
│   ├── main_test.go
│   ├── test.test
│   ├── user_test.go
│   └── users.json
├── main.go
├── compose.yaml
├── dockerfile
├── env
├── go.mod
├── go.sum
└── README.md

```

### Profiling :

  **Benchmark:**
    
    go test -run=XXX -bench . -benchmem

  **CPU Profiling:**
    
    go test -run=XXX -cpuprofile cpu.prof -bench .

    go tool pprof cpu.prof

