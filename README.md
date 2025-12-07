# Go API - Hospital Management System

A Go-based RESTful API for hospital management system built with Clean Architecture principles.

## 🚀 Tech Stack

- Go 1.25.4 + Gin Framework
- PostgreSQL 15 + GORM
- JWT Authentication + Bcrypt
- Clean Architecture (Hexagonal)

## 📋 Prerequisites

- Go 1.25.4+
- Docker (for PostgreSQL)

## ⚙️ Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/Thanamin/go-api.git
cd go-api
```

### 2. Install Go dependencies

```bash
go mod download
```

### 3. Setup environment variables

**For Docker Compose (PostgreSQL):**
```bash
cp docker/docker-compose.env_template docker/.env
```

**For Go API (Project):**
```bash
cp .env.example .env
```

Edit `.env` in project root as needed:
- `DB_HOST=localhost` (connect to Docker PostgreSQL from host)
- `PORT=3000` (or your preferred port)
- `JWT_SECRET` and `HASH_SECRET` (change to secure keys)

### 4. Start PostgreSQL with Docker Compose

```bash
docker compose -f docker/docker-compose.yaml --env-file docker/.env up -d postgres
```

### 5. Run database migrations

```bash
docker exec -i go-api-postgres psql -U postgres -d go_api_db < migrations/001_initial_schema.sql
```

### 6. Start the API server

The `.env` file in project root will be automatically loaded by `make start`:

```bash
make start
```

### 7. Test the API

```bash
curl http://localhost:3000/health
```

Note: Use the PORT from your `.env` file

## 🐳 Docker Database Commands

## 🛠️ Development

**Hot reload (auto-restart on file changes)**
```bash
go install github.com/cespare/reflex@latest
make dev
```

**Build binary**
```bash
make build
./bin/api
```

**Run tests**
```bash
go test ./...              # All tests
go test -cover ./...       # With coverage
go test -v ./...           # Verbose
```

## 📡 API Endpoints

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | /health | No | Check API status |
| POST | /staff/create | No | Create new staff member |
| POST | /staff/login | No | Login and get JWT token |
| GET | /patient/search | Yes | Search patients with filters |

**Authentication:** Use JWT Bearer token from `/staff/login`

**Postman Collection:** See `document/api-postman/` for detailed API documentation

## 📦 Database Schema

- **hospitals**: Hospital information
- **staffs**: Staff members and authentication  
- **patients**: Patient records

See `migrations/001_initial_schema.sql` for complete schema.

## 🚨 Troubleshooting

## 📄 License

This project is private and proprietary.

## 📧 Contact

For questions or support, please contact the development team.

