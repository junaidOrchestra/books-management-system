# Books Management System

## Overview

The **Books Management System** is a RESTful API built using **Golang**, **Gin**, and **GORM** to manage books in a database. It includes features like creating, updating, deleting, and retrieving books. The project also integrates **Redis** for caching and **Kafka** for event streaming.

## Features

- CRUD operations for books
- Redis caching for optimized performance
- Kafka integration for event-driven architecture
- Swagger documentation for API endpoints
- Docker support for easy deployment

## Technologies Used

- **Golang** (Gin, GORM, Uber Fx for dependency injection)
- **SQLite** (for lightweight database management)
- **Redis** (for caching)
- **Kafka** (for event streaming)
- **Swagger** (API documentation)
- **Docker & Docker Compose** (for containerization)

## Setup & Installation

### Prerequisites

- **Go** (1.18 or later)
- **Docker & Docker Compose**
- **Redis**
- **Kafka**

### Clone the Repository

```sh
git clone https://github.com/junaidOrchestra/books-management-system.git
cd books-management-system
```

### Install Dependencies

```sh
go mod tidy
```

### Run the Application

```sh
go run cmd/main.go
```

## API Documentation (Swagger)

Swagger documentation is available at:

```
http://localhost:8080/swagger/index.html
```

To regenerate Swagger docs, run:

```sh
swag init -g cmd/main.go --parseDependency --parseInternal --output ./docs

```

## Running with Docker

The application, along with **Redis** and **Kafka**, can be started using Docker:

```sh
docker-compose up --build
```

To stop the containers:

```sh
docker-compose down
```

## Project Structure

```
books-management-system/
│── cmd/                    # Main entry point
│── internal/
│   ├── controllers/        # API Controllers
│   ├── models/             # Database Models
│   ├── services/           # Business Logic
│   ├── router/             # Gin Router
│── utils/                  # Utility functions
│── docs/                   # Swagger documentation
│── docker-compose.yml      # Docker configuration
│── go.mod                  # Go module file
│── README.md               # Project documentation
```

## Contribution

Feel free to fork this repository and submit a pull request with improvements!

## License

This project is licensed under the MIT License.

