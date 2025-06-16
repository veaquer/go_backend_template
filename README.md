# Go Backend Template

A modern, production-ready Go backend template with built-in authentication, email verification, and caching support. This template provides a solid foundation for building scalable web applications using Go.

## Features

- 🔐 **Authentication System**
  - JWT-based authentication with access and refresh tokens
  - User registration and login
  - Email verification flow
  - Password hashing and security

- 🏗️ **Clean Architecture**
  - Layered architecture (Handler → Service → Repository)
  - Dependency injection using Wire
  - Clear separation of concerns
  - Modular and maintainable code structure

- 🔧 **Technical Stack**
  - [Gin](https://github.com/gin-gonic/gin) - Web framework
  - [GORM](https://gorm.io) - ORM for PostgreSQL
  - [Redis](https://redis.io) - Caching layer
  - [JWT](https://github.com/golang-jwt/jwt) - Authentication tokens
  - [Air](https://github.com/air-verse/air) - Live reload for development
  - [Viper](https://github.com/spf13/viper) - Configuration management
  - [Zap](https://github.com/uber-go/zap) - Logging

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── auth/               # Authentication middleware and token management
│   ├── bootstrap/          # Application bootstrapping and dependency injection
│   ├── cache/             # Redis cache implementation
│   ├── config/            # Configuration management
│   ├── db/                # Database connection and management
│   ├── middleware/        # Custom middleware
│   ├── migration/         # Database migrations
│   ├── user/             # User module (handler, service, repository, model)
│   └── verification/      # Email verification module
├── pkg/
│   ├── email/            # Email service implementation
│   ├── errors/           # Custom error handling
│   ├── hash/             # Password hashing utilities
│   ├── utils/            # Common utilities
│   └── validator/        # Request validation
└── build/
    └── Dockerfile        # Docker configuration
```

## Getting Started

### Prerequisites

- Go 1.24 or higher
- PostgreSQL
- Redis
- Docker (optional)

### Configuration

Create a `.env` file in the root directory with the following variables:

```env
# Server Configuration
ENV=development
BACKEND_URL=http://localhost:4000
PORT=4000

# Database Configuration
DSN=postgres://user:password@localhost:5432/dbname?sslmode=disable

# Redis Configuration
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT Configuration
JWT_ACCESS_SECRET=your_access_secret
JWT_REFRESH_SECRET=your_refresh_secret
JWT_ACCESS_EXPIRATION=15m
JWT_REFRESH_EXPIRATION=720h

# Email Configuration
EMAIL_FROM=your@email.com
EMAIL_HOST=smtp.example.com
EMAIL_PORT=587
EMAIL_USER=your_username
EMAIL_PASS=your_password
```

### Running the Application

#### Using Docker Compose

```bash
# Build and run the application
./scripts/run.sh
```

#### Local Development

```bash
# Install dependencies
go mod download

# Run the application with hot reload
air
```

## API Endpoints

### Authentication

```
POST /auth/register     # Register a new user
POST /auth/login       # Login user
GET  /auth/verify_email # Verify email address
POST /auth/refresh     # Refresh access token
POST /auth/logout      # Logout user
```

### User

```
GET  /user/profile     # Get user profile (requires authentication)
```

## Development Guidelines

### Adding a New Feature

1. Create appropriate models in the relevant module
2. Implement the repository layer for database operations
3. Create the service layer with business logic
4. Add handler methods for HTTP endpoints
5. Register routes in the router
6. Update dependency injection in `bootstrap/wire.go`

### Error Handling

The template uses custom error handling with proper HTTP status codes. Use the `apperror` package for consistent error responses:

```go
if err != nil {
    return apperror.New("Invalid username", http.StatusUnauthorized)
}
```

### Caching

The template includes Redis caching for user profiles. Implement similar caching for other resources as needed:

```go
// Get user from cache first
user, err := redis.GetUserProfile(ctx, userID)
if err == nil {
    return user, nil
}

// If not in cache, get from database and cache it
user, err = repo.GetUserByID(ctx, userID)
redis.SetUserProfile(ctx, user, duration)
```

## Security Features

- Password hashing using modern algorithms
- JWT-based authentication with refresh tokens
- Email verification for new accounts
- Secure session management
- Input validation and sanitization
- Error handling without sensitive information exposure

## Contributing

Feel free to submit issues, fork the repository and create pull requests for any improvements.
