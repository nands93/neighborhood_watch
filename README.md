# Hyper-Local Community Alert System (Temporary Name: Neighborhood Watch)

> ## 🚧 Status: In development 🚧

A work-in-progress platform for community safety: report and visualize neighborhood alerts on an interactive map, with real-time updates and secure user authentication.

---

## ✨ Tech Stack

| Category | Technology |
| :--- | :--- |
| Backend | Go, Gin Framework |
| Database | PostgreSQL with PostGIS |
| Real-Time | WebSockets (Gorilla/Nhooyr) |
| Authentication | JWT with Password Hashing (Argon2) |
| Infrastructure | Docker, Docker Compose |
| Frontend | React (Vite), Leaflet.js (Map) |
| Automation | Makefile |

---

## 🚀 Quick Start

Prerequisites:
- Docker and Docker Compose
- GNU Make (recommended)

1) Configure environment
- Set required env vars via a `.env` file or directly in `docker-compose.yml` (DB credentials, service ports, etc.).

2) Start the stack (recommended)
- `make up` — build and start all services
- `make down` — stop and remove services/volumes
- Check the Makefile for additional targets

Alternative (without make):
- `docker compose up -d`
- `docker compose down -v`

Notes:
- Exposed ports and service URLs are defined in `docker-compose.yml`.
- First startup may initialize the database (including PostGIS) and admin tooling.

---

## 📝 Features & Roadmap

- [x] Go backend skeleton (Gin)
- [x] Full Docker-based dev environment
- [x] Resilient PostgreSQL connection
- [x] User Registration API (`POST /api/v1/users`)
  - [x] Email validation and password strength checks
  - [x] Secure password hashing with Argon2
- [x] Integrated database admin tool (pgAdmin)
- [ ] User Login API (`POST /api/v1/auth/token`)
  - [ ] JWT token generation
- [ ] Authentication middleware for protected routes
- [ ] Alerts API: create and view alerts
- [ ] Interactive map integration on the frontend
- [ ] Real-time notifications via WebSockets

---

## 🏛️ Project Architecture

- Backend: Go (Gin), exposes a REST API, handles auth, validation, and DB access.
- Database: PostgreSQL with PostGIS for geospatial data (alerts, locations).
- Frontend: React (Vite) with Leaflet.js for the interactive map.
- Real-time: WebSockets for live alert updates (planned).
- Infrastructure: Docker + Docker Compose for local development; Makefile for automation.

For deeper details (repository layout, database schema, security choices, etc.), see ARCHITECTURE.md.

---

## ⚙️ API Endpoints

| Method | Path | Description | Auth | Status |
| :---: | :--- | :--- | :---: | :---: |
| POST | /api/v1/users | Register a new user (validates email, enforces password policy, Argon2 hashing) | No | ✅ |
| POST | /api/v1/auth/token | Obtain JWT access token (login) | No | ⏳ |
| POST | /api/v1/alerts | Create a new alert (geolocated) | Yes (JWT) | ⏳ |
| GET | /api/v1/alerts | List or query alerts | Optional (JWT) | ⏳ |
| WS | /api/v1/alerts/stream | Real-time alerts stream (WebSockets) | Yes (JWT) | ⏳ |

Legend: ✅ implemented · ⏳ planned

---

## 🤝 Contact

- Project: Neighborhood Watch
- Issues & Feature Requests: open a GitHub issue in this repository
- Maintainer: Fernando Marques — https://www.linkedin.com/in/nandomarques/ — https://github.com/nands93