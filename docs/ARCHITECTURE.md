# Architecture

Detailed technical documentation for Neighborhood Watch. For a high-level overview and quick start, see [README.md](README.md).

---

## Repository Structure

- Root
  - [.env](.env)
  - [.gitignore](.gitignore)
  - [docker-compose.yml](docker-compose.yml)
  - [Makefile](Makefile)
  - [README.md](README.md)
  - [ARCHITECTURE.md](ARCHITECTURE.md)
- Backend: [backend/](backend)
  - Build: [backend/Dockerfile](backend/Dockerfile)
  - Go modules: [backend/go.mod](backend/go.mod), [backend/go.sum](backend/go.sum)
  - Entry point: [backend/cmd/server/main.go](backend/cmd/server/main.go)
  - Database init SQL: [backend/database/init.sql](backend/database/init.sql)
  - Internal packages: [backend/internal/](backend/internal)
    - Auth: [backend/internal/auth](backend/internal/auth)
      - [`auth.GenerateArgon2Hash`](backend/internal/auth/password.go)
      - [`auth.ValidateEmail`](backend/internal/auth/validate_email.go)
      - [`auth.ValidateEmailStrict`](backend/internal/auth/validate_email.go)
      - [`auth.ValidatePasswordStrength`](backend/internal/auth/validate_password.go)
      - [`auth.DefaultPasswordStrength`](backend/internal/auth/validate_password.go)
    - Database: [backend/internal/database](backend/internal/database)
      - [`database.ConnectToDB`](backend/internal/database/db.go)
    - Handler: [backend/internal/handler](backend/internal/handler)
      - [`handler.HealthCheck`](backend/internal/handler/health_handler.go)
      - [`handler.RegisterUser`](backend/internal/handler/user_handler.go)
    - Model: [backend/internal/model](backend/internal/model)
      - [`model.User`](backend/internal/model/user_struct.go)
    - Repository: [backend/internal/repository](backend/internal/repository)
      - [`repository.CreateUser`](backend/internal/repository/user_repository.go)
- Frontend scaffold: [frontend/](frontend)

---

## Backend

- Module: vizinhanca ([backend/go.mod](backend/go.mod))
- Server bootstrap: [backend/cmd/server/main.go](backend/cmd/server/main.go)
  - DB connection on startup via [`database.ConnectToDB`](backend/internal/database/db.go)
  - Routes (Gin):
    - GET /api/v1/health → [`handler.HealthCheck`](backend/internal/handler/health_handler.go)
    - POST /api/v1/users → [`handler.RegisterUser`](backend/internal/handler/user_handler.go)
- Handlers:
  - [`handler.RegisterUser`](backend/internal/handler/user_handler.go)
    - Validates email via [`auth.ValidateEmail`](backend/internal/auth/validate_email.go)
    - Enforces password policy via [`auth.DefaultPasswordStrength`](backend/internal/auth/validate_password.go) and [`auth.ValidatePasswordStrength`](backend/internal/auth/validate_password.go)
    - Hashes password via [`auth.GenerateArgon2Hash`](backend/internal/auth/password.go)
    - Persists user via [`repository.CreateUser`](backend/internal/repository/user_repository.go)
  - [`handler.HealthCheck`](backend/internal/handler/health_handler.go) returns simple service status
- Build image: [backend/Dockerfile](backend/Dockerfile)

---

## Database

- Engine: PostgreSQL with PostGIS (Docker image: postgis/postgis) configured in [docker-compose.yml](docker-compose.yml)
- Schema bootstrap: [backend/database/init.sql](backend/database/init.sql)
  - users table: id, username, email, password_hash, created_at, updated_at, last_login
- Connection pooling:
  - [`database.ConnectToDB`](backend/internal/database/db.go) builds DSN from [.env](.env), initializes `pgxpool.Pool`, retries with backoff and health check (Ping)
- Admin tool: pgAdmin service defined in [docker-compose.yml](docker-compose.yml)

---

## Code Highlights

- Data model:
  - [`model.User`](backend/internal/model/user_struct.go)
- Auth utilities:
  - Password hashing: [`auth.GenerateArgon2Hash`](backend/internal/auth/password.go) (Argon2id, random salt, base64 encoding)
  - Email validation: [`auth.ValidateEmail`](backend/internal/auth/validate_email.go), strict checks [`auth.ValidateEmailStrict`](backend/internal/auth/validate_email.go)
  - Password policy: [`auth.DefaultPasswordStrength`](backend/internal/auth/validate_password.go) + [`auth.ValidatePasswordStrength`](backend/internal/auth/validate_password.go)
- HTTP layer:
  - Health: [`handler.HealthCheck`](backend/internal/handler/health_handler.go)
  - Registration: [`handler.RegisterUser`](backend/internal/handler/user_handler.go)
- Persistence:
  - Create user: [`repository.CreateUser`](backend/internal/repository/user_repository.go) using [`database.DB`](backend/internal/database/db.go) pool

---

## Configuration

- Environment variables (see [.env](.env)):
  - POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_HOST, POSTGRES_PORT
  - PGADMIN_DEFAULT_EMAIL, PGADMIN_DEFAULT_PASSWORD
- Services and ports (see [docker-compose.yml](docker-compose.yml)):
  - Backend: 8080/tcp
  - PostgreSQL: 5432/tcp
  - pgAdmin: 5050/tcp
- Build/runtime:
  - Backend image built via [backend/Dockerfile](backend/Dockerfile)
  - DSN assembled in [`database.ConnectToDB`](backend/internal/database/db.go)

---

## Development Notes

- Orchestration: [Makefile](Makefile)
  - `make up` / `make down` / `make re` / `make clean` / `make fclean`
- Compose: `docker compose up -d` and `docker compose down -v` (see [docker-compose.yml](docker-compose.yml))
- Logs:
  - Backend: `docker logs -f backend`
  - PostGIS: `docker logs -f postgis`
  - pgAdmin: `docker logs -f pgadmin`
- Hot reload: not configured (consider air for local DX)
- Module name: vizinhanca (import paths use `vizinhanca/...`)

---

## Troubleshooting

- Backend cannot reach DB:
  - Ensure db is healthy (healthcheck in [docker-compose.yml](docker-compose.yml))
  - Verify env vars in [.env](.env)
  - Connection uses retries in [`database.ConnectToDB`](backend/internal/database/db.go)
- pgAdmin cannot connect:
  - Use host `db`, port `5432`, credentials from [.env](.env)
  - Ensure the `postgis` container is running
- Reinitialize DB:
  - `make down` (removes volumes) then `make up` to rerun [backend/database/init.sql](backend/database/init.sql)

---

## Security

- Passwords:
  - Hashing with Argon2id via [`auth.GenerateArgon2Hash`](backend/internal/auth/password.go)
  - Note: avoid logging raw passwords. Remove debug logging in [`auth.ValidatePasswordStrength`](backend/internal/auth/validate_password.go) before production.
- Validation:
  - Email format checks: [`auth.ValidateEmail`](backend/internal/auth/validate_email.go), [`auth.ValidateEmailStrict`](backend/internal/auth/validate_email.go)
  - Password policy: [`auth.DefaultPasswordStrength`](backend/internal/auth/validate_password.go)
- Authentication:
  - JWT issuance and middleware planned; not yet implemented