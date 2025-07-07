# URL Shortening Service

A high-performance URL shortening service built with Go backend and React frontend, featuring user authentication, secure password hashing, rate limiting, and caching for optimal performance.

## 🚀 Features

- **URL Shortening**: Convert long URLs into short, manageable links with unique slugs
- **User Authentication**: JWT-based authentication with secure login/registration
- **URL Management**: List and manage all your shortened URLs
- **Modern Frontend**: React with TypeScript, Vite, and Tailwind CSS
- **Dark Theme**: Beautiful dark-themed user interface
- **Password Security**: Secure password hashing using bcrypt
- **Rate Limiting**: Built-in rate limiting to prevent abuse
- **Redis Caching**: Fast URL resolution with Redis caching
- **Database Persistence**: PostgreSQL for reliable data storage
- **Clean Architecture**: Well-structured codebase following Go best practices
- **Docker Support**: Easy deployment with Docker and Docker Compose
- **Database Migration**: Automated database schema management
- **Unique Constraints**: Prevents duplicate URL shortenings per user
- **Copy to Clipboard**: Easy URL sharing with one-click copy functionality

## 🛠️ Technology Stack

### Backend

- **Go**: 1.24.4 with Fiber framework
- **Database**: PostgreSQL 17.3
- **Cache**: Redis
- **Authentication**: JWT (JSON Web Tokens) with HTTPOnly cookies
- **Password Hashing**: bcrypt
- **Validation**: Go Playground Validator
- **ORM**: GORM with PostgreSQL driver
- **Environment Management**: Godotenv

### Frontend

- **React**: 18.3.1 with TypeScript
- **Build Tool**: Vite
- **Styling**: Tailwind CSS with dark theme
- **Router**: React Router DOM
- **HTTP Client**: Axios with credentials support
- **State Management**: React Context API

### Infrastructure

- **Containerization**: Docker & Docker Compose
- **CORS**: Configured for cross-origin requests

## 📁 Project Structure

```
URL_shortening/
├── cmd/                        # Application entrypoint
│   └── main.go
├── front/                      # React frontend
│   ├── public/                 # Static assets
│   ├── src/
│   │   ├── components/         # React components
│   │   │   └── ProtectedRoute.tsx
│   │   ├── contexts/           # React contexts
│   │   │   └── AuthContext.tsx
│   │   ├── pages/              # Application pages
│   │   │   ├── Dashboard.tsx   # URL shortening page
│   │   │   ├── Login.tsx       # Login page
│   │   │   ├── Register.tsx    # Registration page
│   │   │   └── MyUrls.tsx      # URL management page
│   │   ├── services/           # API services
│   │   │   └── api.ts
│   │   ├── types/              # TypeScript types
│   │   │   └── index.ts
│   │   └── App.tsx
│   ├── package.json
│   ├── vite.config.ts
│   └── tailwind.config.js
├── infra/                      # Infrastructure layer
│   ├── config/
│   │   └── environment/
│   │       └── config.go
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
│       │   ├── login.go
│       │   ├── register.go
│       │   ├── logout.go
│       │   └── me.go
│       └── urlShortening/      # URL shortening use cases
│           ├── urlShortening_useCase.go
│           └── list.go         # List user URLs
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

# Frontend Configuration
FRONTEND_URL=http://localhost:3000
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

3. **Build and run the backend**

   ```bash
   make build
   make run
   ```

4. **Install and run the frontend**

   ```bash
   cd front
   npm install
   npm run dev
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

3. **Run the backend**

   ```bash
   make run
   ```

4. **Install and run the frontend**
   ```bash
   cd front
   npm install
   npm run dev
   ```

The backend will be available at `http://localhost:8181` and the frontend at `http://localhost:3000`

## 🖥️ Frontend Application

### Available Pages

1. **Login Page** (`/login`)

   - User authentication with email and password
   - Redirect to dashboard after successful login

2. **Registration Page** (`/register`)

   - New user registration with name, email, and password
   - Password confirmation validation

3. **Dashboard** (`/dashboard`)

   - Main URL shortening interface
   - Form to input URLs for shortening
   - Display of shortened URL result
   - Copy-to-clipboard functionality

4. **My URLs** (`/my-urls`)
   - List of all user's shortened URLs
   - Display original and shortened URLs
   - Creation date and time information
   - Copy-to-clipboard functionality for each URL

### Navigation

- **Dashboard**: Main page for creating new shortened URLs
- **My URLs**: Access via "Minhas URLs" button in header
- **Logout**: Available from both Dashboard and My URLs pages

### Features

- **Dark Theme**: Consistent dark UI theme across all pages
- **Responsive Design**: Mobile-friendly interface
- **Protected Routes**: Authentication required for Dashboard and My URLs
- **Real-time Feedback**: Success/error messages for user actions
- **Copy to Clipboard**: One-click copying of shortened URLs

## 📋 How to Use

### Getting Started

1. **Start the Services**

   ```bash
   # Terminal 1 - Start infrastructure
   docker-compose up postgres redis -d

   # Terminal 2 - Start backend
   make run

   # Terminal 3 - Start frontend
   cd front && npm run dev
   ```

2. **Access the Application**

   - Open your browser and go to `http://localhost:3000`
   - You'll be redirected to the login page

3. **Create an Account**

   - Click "Não tem conta? Registre-se"
   - Fill in your name, email, and password
   - Click "Registrar"

4. **Login**

   - Use your email and password to login
   - You'll be redirected to the dashboard

5. **Shorten URLs**

   - Enter a URL in the form (e.g., `https://www.google.com`)
   - Click "Encurtar URL"
   - Copy the shortened URL using the "Copiar" button

6. **Manage Your URLs**
   - Click "Minhas URLs" in the header
   - View all your shortened URLs
   - Copy any URL using the "Copiar" button
   - See creation dates and times

### Testing the Application

- **Public Access**: Visit `http://localhost:8181/{slug}` (e.g., `http://localhost:8181/abc12345`) to test URL redirection
- **API Testing**: Use tools like Postman or curl to test API endpoints
- **Frontend Testing**: All functionality is accessible through the web interface

## 🔧 Troubleshooting

### Common Issues

1. **CORS Errors**

   - Ensure `FRONTEND_URL=http://localhost:3000` is set in your `.env` file
   - Verify the frontend is running on port 3000

2. **Database Connection Issues**

   - Check if PostgreSQL is running: `docker-compose ps`
   - Verify database credentials in `.env` file
   - Ensure database exists: `url_shortening`

3. **Redis Connection Issues**

   - Check if Redis is running: `docker-compose ps`
   - Verify Redis address in `.env` file: `REDIS_ADDRESS=localhost:6379`

4. **JWT Authentication Issues**

   - Clear browser cookies and try again
   - Check if `JWT_SECRET` is set in `.env` file
   - Verify middleware is properly applied to protected routes

5. **Frontend Build Issues**

   - Delete `node_modules` and `package-lock.json`
   - Run `npm install` again
   - Ensure Node.js version is 18+ and npm is up to date

6. **URL Shortening Not Working**
   - Check if URL is valid (must include `http://` or `https://`)
   - Verify user is authenticated
   - Check backend logs for errors

### Debug Steps

1. **Check Service Status**

   ```bash
   docker-compose ps
   ```

2. **View Logs**

   ```bash
   # Backend logs
   docker-compose logs url-shortener

   # Database logs
   docker-compose logs postgres

   # Redis logs
   docker-compose logs redis
   ```

3. **Test Database Connection**

   ```bash
   psql -h localhost -U root -d url_shortening
   ```

4. **Test API Endpoints**

   ```bash
   # Health check
   curl http://localhost:8181/

   # List URLs (replace with actual token)
   curl -X GET http://localhost:8181/urls -H "Cookie: token=your-jwt-token"
   ```

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

#### Get Current User (Protected)

```http
GET /auth/me
Cookie: token=<jwt-token>
```

**Response:**

```json
{
  "user": {
    "id": "user-id",
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

#### Logout User (Protected)

```http
POST /auth/logout
Cookie: token=<jwt-token>
```

**Response:**

```json
{
  "message": "Logout successful"
}
```

### URL Shortening Endpoints

#### Shorten URL (Protected)

```http
POST /register
Content-Type: application/json
Cookie: token=<jwt-token>

{
  "url": "https://example.com/very-long-url-that-needs-shortening"
}
```

**Response:**

```json
{
  "shortUrl": "http://localhost:8181/abc12345",
  "originalUrl": "https://example.com/very-long-url-that-needs-shortening"
}
```

#### List User URLs (Protected)

```http
GET /urls
Cookie: token=<jwt-token>
```

**Response:**

```json
{
  "urls": [
    {
      "ID": "url-id",
      "UrlOriginal": "https://example.com/very-long-url",
      "UrlShortened": "http://localhost:8181/abc12345",
      "Slug": "abc12345",
      "CreatedAt": "2024-01-01T12:00:00Z"
    }
  ]
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

- `POST /register` - Create shortened URLs
- `GET /urls` - List user's shortened URLs
- `GET /auth/me` - Get current user information
- `POST /auth/logout` - Logout user

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
- **URL registration (`/register`)**: 100 requests per minute
- **URL listing (`/urls`)**: 50 requests per minute

## 📈 Performance Features

- **Redis Caching**: Shortened URLs are cached for 3 minutes for faster resolution
- **Database Indexing**: Optimized queries with proper indexing
- **Connection Pooling**: Efficient database connection management
- **Unique Constraints**: Prevents duplicate URL shortenings per user and ensures unique slugs

## 🛡️ Security Features

- **JWT Authentication**: Secure token-based authentication with HTTPOnly cookies
- **Password Security**: Secure password hashing using bcrypt
- **Input Validation**: Comprehensive request validation
- **CORS Protection**: Built-in CORS middleware configured for frontend
- **Rate Limiting**: Protection against brute force attacks
- **Unique Constraints**: Database-level constraint preventing duplicate URLs per user
- **Protected Routes**: Frontend route protection with authentication guards

## 🔧 Development Commands

### Backend Commands

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

### Frontend Commands

```bash
# Navigate to frontend directory
cd front

# Install dependencies
npm install

# Run development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
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
- [ ] URL editing and deletion
- [ ] Export URL data
- [ ] Mobile responsive improvements
- [ ] PWA (Progressive Web App) support

## 📞 Support

For support, please open an issue in the GitHub repository or contact the development team.

## 🎯 Key Improvements Implemented

- ✅ **Password Security**: Implemented bcrypt password hashing
- ✅ **Database Constraints**: Added unique constraint for user-URL combinations
- ✅ **Slug System**: Implemented unique slug generation for URLs
- ✅ **Security Headers**: Enhanced security with proper authentication middleware
- ✅ **Error Handling**: Comprehensive error handling throughout the application
- ✅ **Modern Frontend**: React with TypeScript, Vite, and Tailwind CSS
- ✅ **URL Management**: Complete URL listing and management system
- ✅ **Authentication System**: JWT with HTTPOnly cookies for security
- ✅ **Dark Theme**: Beautiful dark-themed user interface
- ✅ **CORS Configuration**: Proper cross-origin resource sharing setup
- ✅ **Protected Routes**: Frontend route protection with authentication guards
