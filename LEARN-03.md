Certainly! For a web API project with a database connection in Go, you can follow a structured layout that promotes clean architecture and scalability. Below is an example project structure with explanations and some sample code.

# PART 1

### Example Project Structure

```
yourproject/
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── database.go
│   ├── handlers/
│   │   └── user.go
│   ├── models/
│   │   └── user.go
│   ├── repositories/
│   │   └── user_repository.go
│   └── router/
│       └── router.go
└── pkg/
    └── utils/
        └── utils.go
```

### Explanation of Each Directory and File

1. **`go.mod` and `go.sum`**:
   - Manage module dependencies.

2. **`main.go`**:
   - Entry point of the application that initializes the server.

3. **`cmd/server/main.go`**:
   - Initializes and starts the HTTP server.

4. **`internal/config/config.go`**:
   - Configuration handling, e.g., reading from environment variables or configuration files.

5. **`internal/database/database.go`**:
   - Database connection setup and management.

6. **`internal/handlers/user.go`**:
   - HTTP handlers for user-related API endpoints.

7. **`internal/models/user.go`**:
   - Data models representing the structure of database tables.

8. **`internal/repositories/user_repository.go`**:
   - Data access layer, interacting with the database.

9. **`internal/router/router.go`**:
   - Setting up routes and middleware.

10. **`pkg/utils/utils.go`**:
    - Utility functions that can be shared across the project.

### Detailed Example

#### `go.mod`

```go
module github.com/yourusername/yourproject

go 1.19

require (
    github.com/gin-gonic/gin v1.7.7
    github.com/jmoiron/sqlx v1.3.4
    github.com/joho/godotenv v1.4.0
)
```

#### `main.go`

```go
package main

import (
    "github.com/yourusername/yourproject/cmd/server"
)

func main() {
    server.Run()
}
```

#### `cmd/server/main.go`

```go
package server

import (
    "log"
    "github.com/yourusername/yourproject/internal/config"
    "github.com/yourusername/yourproject/internal/database"
    "github.com/yourusername/yourproject/internal/router"
    "github.com/gin-gonic/gin"
)

func Run() {
    config.LoadConfig()
    db := database.InitDB()
    r := gin.Default()
    router.SetupRoutes(r, db)
    if err := r.Run(); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
```

#### `internal/config/config.go`

```go
package config

import (
    "log"
    "github.com/joho/godotenv"
)

func LoadConfig() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }
}
```

Setup environment variable in `.env` file of the project.

#### `internal/database/database.go`

```go
package database

import (
    "log"
    "os"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func InitDB() *sqlx.DB {
    dsn := os.Getenv("DATABASE_URL")
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    return db
}
```

At the environment `.env` to add this **DSN** value:
```
DATABASE_URL=postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable

```

#### `internal/handlers/user.go`

```go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/yourusername/yourproject/internal/repositories"
    "github.com/jmoiron/sqlx"
)

func GetUsers(db *sqlx.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        users, err := repositories.GetAllUsers(db)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, users)
    }
}
```

#### `internal/models/user.go`

```go
package models

type User struct {
    ID    int    `db:"id" json:"id"`
    Name  string `db:"name" json:"name"`
    Email string `db:"email" json:"email"`
}
```

Match to Postgres SQL/table creation as:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);

```

#### `internal/repositories/user_repository.go`

```go
package repositories

import (
    "github.com/yourusername/yourproject/internal/models"
    "github.com/jmoiron/sqlx"
)

func GetAllUsers(db *sqlx.DB) ([]models.User, error) {
    users := []models.User{}
    err := db.Select(&users, "SELECT id, name, email FROM users")
    return users, err
}
```

#### `internal/router/router.go`

```go
package router

import (
    "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"
    "github.com/yourusername/yourproject/internal/handlers"
)

func SetupRoutes(r *gin.Engine, db *sqlx.DB) {
    r.GET("/users", handlers.GetUsers(db))
}
```

#### `pkg/utils/utils.go`

```go
package utils

import (
    "log"
    "github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, message string) {
    log.Printf("%s: %v", message, err)
    c.JSON(500, gin.H{"error": message})
}
```

### Summary

This structure helps in keeping your project organized by separating concerns:

- **Configuration**: Load environment variables and configuration settings.
- **Database**: Manage database connections.
- **Handlers**: Define HTTP handlers that process incoming requests.
- **Models**: Define the structure of your data.
- **Repositories**: Interact with the database to perform CRUD operations.
- **Router**: Set up routing and middleware for the HTTP server.
- **Utilities**: Common utility functions used throughout the project.

This structure is flexible and can be adjusted as your project evolves. It promotes maintainability and scalability, making it easier to manage a growing codebase.

# PART 2
