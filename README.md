# jlGoServer

A simple Go server with SQLite database integration using GORM.

## Features

- RESTful API for user management
- SQLite database for data persistence
- Static file serving
- Form handling

## Project Structure

```
jlGoServer/
├── db/               # Database connection logic
├── handlers/         # HTTP request handlers
├── models/           # Data models
├── static/           # Static files
│   ├── form.html     # Simple form page
│   ├── index.html    # Homepage
│   └── users.html    # User management UI
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
├── jlgo.db           # SQLite database file (created at runtime)
├── README.md         # Project documentation
└── server.go         # Main server application
```

## API Endpoints

### User Management

- `GET /api/users` - Get all users
- `GET /api/users/{id}` - Get a specific user
- `POST /api/users` - Create a new user
- `PUT /api/users/{id}` - Update a user
- `DELETE /api/users/{id}` - Delete a user

## User Interface

- `/users.html` - Web interface for user management
- `/form.html` - Simple form submission
- `/hello` - Basic hello endpoint

## Getting Started

1. Make sure Go is installed on your system
2. Clone the repository
3. Run the server:
   ```
   go run server.go
   ```
4. Access the web interface at http://localhost:8080/users.html

## Database Schema

### User

| Field     | Type      | Description                     |
| --------- | --------- | ------------------------------- |
| ID        | uint      | Primary key                     |
| Name      | string    | User's name                     |
| Email     | string    | User's email (unique)           |
| Address   | string    | User's address                  |
| JoinedAt  | time.Time | When user joined                |
| CreatedAt | time.Time | Record creation time            |
| UpdatedAt | time.Time | Last update time                |
| DeletedAt | time.Time | Deletion time (for soft delete) |

## Future Improvements

- User authentication
- More comprehensive API validation
- Additional database tables and relationships
- Structured logging
- Configuration management
- Containerization
