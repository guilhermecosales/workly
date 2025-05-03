# Workly

> Workly is a simple job opportunities API written in Go, leveraging the Gin web framework, GORM ORM, and PostgreSQL for
> data persistence. It provides basic CRUD operations to manage job listings, making it easy to list, create, update,
> and
> delete job postings.

---

## Table of Contents

* [Features](#features)
* [Tech Stack](#tech-stack)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Configuration](#configuration)
* [Database Setup](#database-setup)
* [Running the Application](#running-the-application)
* [API Endpoints](#api-endpoints)
* [Project Structure](#project-structure)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)

---

## Features

* List all job opportunities
* Retrieve a single job by ID
* Create new job listings
* Update existing job listings
* Delete job listings
* Input validation and error handling
* Docker support via `docker-compose`

## Tech Stack

* **Language:** Go
* **Web Framework:** Gin
* **ORM:** GORM
* **Database:** PostgreSQL
* **Dependency Management:** Go Modules
* **Containerization:** Docker & Docker Compose

## Prerequisites

* Go 1.18 or higher installed
* Docker & Docker Compose (optional, for running PostgreSQL locally)
* PostgreSQL database (if not using Docker)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/guilhermecosales/workly.git
   cd workly
   ```
2. Download dependencies:

   ```bash
   go mod download
   ```

## Configuration

Create a `.env` file in the project root (or set environment variables) with the following values:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=workly
```

## Database Setup

### Using Docker Compose

If you have Docker installed, you can start a PostgreSQL container:

```bash
docker-compose up -d
```

This will launch PostgreSQL on `localhost:5432` with the credentials configured in `docker-compose.yml`.

### Manual Setup

1. Install PostgreSQL.
2. Create a new database:

   ```sql
   CREATE DATABASE workly;
   ```
3. Ensure your `.env` matches the database credentials.

## Running the Application

Start the API:

```bash
go run main.go
```

By default, the server will run on `http://localhost:8080`.

---

## API Endpoints

| Method | Endpoint               | Description                    |
|--------|------------------------|--------------------------------|
| GET    | `/openings`            | List all jobs                  |
| GET    | `/openings/:openingId` | Get a single job by its ID     |
| POST   | `/openings`            | Create a new job listing       |
| PUT    | `/openings/:openingId` | Update an existing job listing |
| DELETE | `/openings/:openingId` | Delete a job listing           |

### Example Requests

* List jobs:

  ```bash
  curl http://localhost:8080/openings
  ```

* Create a job:

  ```bash
  curl -X POST http://localhost:8080/openings \
    -H "Content-Type: application/json" \
    -d '{
      "title": "Software Engineer",
      "description": "Develop and maintain APIs",
      "company": "TechCorp",
      "location": "Remote",
      "salary": 75000
    }'
  ```

* Update a job:

  ```bash
  curl -X PUT http://localhost:8080/openings/1 \
    -H "Content-Type: application/json" \
    -d '{
      "title": "Senior Software Engineer",
      "salary": 90000
    }'
  ```

* Delete a job:

  ```bash
  curl -X DELETE http://localhost:8080/openings/1
  ```

---

## Project Structure

```
workly/
├── config/           # Environment and database configuration
├── handler/          # HTTP handler functions
├── router/           # Route definitions
├── schemas/          # Data models and validation schemas
├── .gitignore        # Files to ignore in Git
├── docker-compose.yml# Docker Compose file for PostgreSQL
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
└── main.go           # Application entry point
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
