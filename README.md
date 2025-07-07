# URL Shortening Service

A high-performance URL shortening service built with Go, featuring user authentication, secure password hashing, rate limiting, and caching for optimal performance.

## 🚀 Features

- **URL Shortening**: Convert long URLs into short, manageable links with unique slugs
- **User Authentication**: JWT-based authentication with secure login/registration
- **Password Security**: Secure password hashing using bcrypt
- **Rate Limiting**: Built-in rate limiting to prevent abuse
- **Redis Caching**: Fast URL resolution with Redis caching
- **Database Persistence**: PostgreSQL for reliable data storage
- **Clean Architecture**: Well-structured codebase following Go best practices
- **Docker Support**: Easy deployment with Docker and Docker Compose
- **Database Migration**: Automated database schema management
- **Unique Constraints**: Prevents duplicate URL shortenings per user

## 🛠️ Technology Stack

- **Backend**: Go 1.24.4 with Fiber framework
- **Database**: PostgreSQL 17.3
- **Cache**: Redis
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Validation**: Go Playground Validator
- **ORM**: GORM with PostgreSQL driver
- **Environment Management**: Godotenv
- **Containerization**: Docker & Docker Compose

## 📁 Project Structure

```
URL_shortening/
├── cmd/                        # Application entrypoint
│   └── main.go
├── config/                     # Configuration management
│   └── environment/
│       └── config.go
├── infra/                      # Infrastructure layer
│   └── db/
│       ├── postgres/           # PostgreSQL implementation
│       │   ├── migration/      # Database migrations
│       │   └── postgres.go
│       └── redis/              # Redis implementation
│           └── redis.go
├── internal/                   # Internal application logic
│   ├── delivery/
│   │   └── httpserver/         # HTTP server implementation
│   │       ├── middleware/     # Authentication middleware
│   │       └── server.go
│   ├── domain/
│   │   └── repository/         # Repository interfaces
│   │       ├── urlShortening_repo/
│   │       └── user_repo/
│   └── useCase/                # Business logic
│       ├── auth/               # Authentication use cases
│       └── urlShortening/      # URL shortening use cases
├── pkg/                        # Shared packages
│   ├── cryptPkg/               # Password encryption utilities
│   ├── env/                    # Environment utilities
│   ├── jwtpkg/                 # JWT utilities
│   └── projectError/           # Custom error handling
├── docker-compose.yml          # Docker services configuration
├── dockerfile                  # Application container
├── Makefile                    # Build automation
└── README.md
```

## 🔧 Installation & Setup

### Prerequisites

- Go 1.24.4 or higher
- Docker and Docker Compose (for containerized setup)
- PostgreSQL 17.3+ (if running locally)
- Redis (if running locally)

### Environment Variables

Create a `.env` file in the project root with the following variables:

```env
# Server Configuration
URL=localhost
PORT=8181

# Database Configuration
DB_DATA_SOURCE=postgres://root:root@localhost:5432/url_shortening?sslmode=disable

# URL Shortening Configuration
URL_SHORTENED_PREFIX=http://localhost:8181

# Redis Configuration
REDIS_ADDRESS=localhost:6379

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key
```

### 🐳 Docker Setup (Recommended)

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd URL_shortening
   ```

2. **Start the infrastructure services**

   ```bash
   docker-compose up postgres redis -d
   ```

3. **Build and run the application**

   ```bash
   make build
   make run
   ```

   _Note: The application container service is currently commented out in docker-compose.yml. You can uncomment it to run the full stack with Docker._

### 🏠 Local Development Setup

1. **Install dependencies**

   ```bash
   go mod tidy
   ```

2. **Start PostgreSQL and Redis**

   ```bash
   docker-compose up postgres redis -d
   ```

3. **Run the application**
   ```bash
   make run
   ```

The application will be available at `http://localhost:8181`

## 📚 API Documentation

### Authentication Endpoints

#### Register User

```http
POST /auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "message": "User registered successfully"
}
```

#### Login User

```http
POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "message": "Login successful"
}
```

### URL Shortening Endpoints

#### Shorten URL (Protected)

```http
POST /url/register
Content-Type: application/json
Cookie: token=<jwt-token>

{
  "url": "https://example.com/very-long-url-that-needs-shortening"
}
```

**Response:**

```json
{
  "message": "http://localhost:8181/abc12345"
}
```

#### Access Shortened URL

```http
GET /:urlShortened
```

**Response:** HTTP 302 Redirect to original URL

#### Health Check

```http
GET /
```

**Response:** "salve! 🤙"

## 🔐 Authentication

The API uses JWT-based authentication with HTTP-only cookies for security. After successful login/registration, a JWT token is set as a cookie and required for accessing protected endpoints.

### Protected Endpoints

- `POST /url/register` - Create shortened URLs

## 🏗️ Database Schema

### Users Table

```sql
CREATE TABLE users (
  id varchar(255) PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL UNIQUE,
  password varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);
```

### URL Shortening Table

```sql
CREATE TABLE url_shortening (
  id varchar(255) PRIMARY KEY,
  id_user varchar(255) NOT NULL,
  url_original varchar(255) NOT NULL,
  url_shortened varchar(255) NOT NULL UNIQUE,
  slug varchar(255) NOT NULL UNIQUE,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),

  FOREIGN KEY (id_user) REFERENCES users(id),
  CONSTRAINT id_user_url_original_unique UNIQUE (id_user, url_original)
);
```

## 🚦 Rate Limiting

The application implements rate limiting to prevent abuse:

- **Authentication endpoints (`/auth/*`)**: 20 requests per minute
- **URL registration (`/url/register`)**: 100 requests per minute

## 📈 Performance Features

- **Redis Caching**: Shortened URLs are cached for 3 minutes for faster resolution
- **Database Indexing**: Optimized queries with proper indexing
- **Connection Pooling**: Efficient database connection management
- **Unique Constraints**: Prevents duplicate URL shortenings per user and ensures unique slugs

## 🛡️ Security Features

- **JWT Authentication**: Secure token-based authentication
- **Password Security**: Secure password hashing using bcrypt
- **Input Validation**: Comprehensive request validation
- **CORS Protection**: Built-in CORS middleware
- **Rate Limiting**: Protection against brute force attacks
- **Unique Constraints**: Database-level constraint preventing duplicate URLs per user

## 🔧 Development Commands

```bash
# Build the application
make build

# Run the application
make run

# Tidy Go modules
make gomod

# Run infrastructure services only
docker-compose up postgres redis -d

# Run with Docker (uncomment url-shortener service first)
# docker-compose up --build
```

## 🚀 Deployment

### Docker Deployment

1. **Uncomment the url-shortener service in docker-compose.yml**
2. **Build and start all services**

   ```bash
   docker-compose up --build
   ```

### Manual Deployment

1. Set up PostgreSQL and Redis instances
2. Configure environment variables
3. Build the application: `make build`
4. Run the binary: `./bin/main`

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/new-feature`
3. Commit your changes: `git commit -am 'Add new feature'`
4. Push to the branch: `git push origin feature/new-feature`
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔮 Future Enhancements

- [ ] Custom short URL aliases
- [ ] Click analytics and statistics
- [ ] URL expiration dates
- [ ] Batch URL shortening
- [ ] API key authentication
- [ ] Admin dashboard
- [ ] URL validation and safety checks
- [ ] QR code generation for shortened URLs
- [ ] URL preview functionality

## 📞 Support

For support, please open an issue in the GitHub repository or contact the development team.

## 🎯 Key Improvements Implemented

- ✅ **Password Security**: Implemented bcrypt password hashing
- ✅ **Database Constraints**: Added unique constraint for user-URL combinations
- ✅ **Slug System**: Implemented unique slug generation for URLs
- ✅ **Security Headers**: Enhanced security with proper authentication middleware
- ✅ **Error Handling**: Comprehensive error handling throughout the application
