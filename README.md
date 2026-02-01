<p align="center">
  <img src="assets/voca_logo.png" alt="VOCA Logo" width="400">
</p>

<h1 align="center">VOCA</h1>

<p align="center">
  <strong>Unlock Your Path</strong><br>
  A modern personality assessment helping young people discover careers that match who they are
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Nuxt-4.3-00DC82?style=flat-square&logo=nuxtdotjs&logoColor=white" alt="Nuxt 4.3">
  <img src="https://img.shields.io/badge/Vue-3-4FC08D?style=flat-square&logo=vuedotjs&logoColor=white" alt="Vue 3">
  <img src="https://img.shields.io/badge/Go-1.23-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go 1.23">
  <img src="https://img.shields.io/badge/Tailwind-4-06B6D4?style=flat-square&logo=tailwindcss&logoColor=white" alt="Tailwind 4">
  <img src="https://img.shields.io/badge/TypeScript-5.9-3178C6?style=flat-square&logo=typescript&logoColor=white" alt="TypeScript 5.9">
  <img src="https://img.shields.io/badge/License-MIT-yellow?style=flat-square" alt="MIT License">
</p>

---

## Overview

VOCA is a web application for administering the **IPIP-50** (International Personality Item Pool) personality assessment based on the Big Five model. Designed with young people (ages 15-18) in mind, it provides an engaging, friendly experience that helps users discover career paths aligned with their personality traits.

## Features

| Feature | Description |
|---------|-------------|
| **Step-by-Step Questionnaire** | Intuitive card-based flow with animated progress tracking |
| **Big Five Assessment** | Measures Extraversion, Agreeableness, Conscientiousness, Emotional Stability, and Openness |
| **Visual Results** | Radar charts and trait cards with personalized score interpretations |
| **Responsive Design** | Beautiful experience on desktop, tablet, and mobile |
| **Multi-Language** | Supports 10 languages for global accessibility |
| **Career Insights** | AI-powered career recommendations based on personality profile |

## Scientific Foundation

This assessment uses the **IPIP-50**, a validated 50-item scale from the International Personality Item Pool. IPIP is a public-domain collection of personality assessment items developed by leading psychology researchers. Unlike proprietary tests, IPIP instruments are freely available for research and practical use.

Learn more at [ipip.ori.org](https://ipip.ori.org/)

## Project Structure

```
voca/
├── frontend/              # Nuxt 4 application
│   ├── app/
│   │   ├── components/    # Vue components
│   │   ├── composables/   # Reusable logic
│   │   └── pages/         # Route pages
│   └── ...
├── backend/               # Go API server
│   ├── cmd/api/           # Application entry point
│   ├── internal/
│   │   ├── config/        # Configuration
│   │   ├── domain/        # Domain models
│   │   ├── handler/       # HTTP handlers
│   │   ├── repository/    # Data access
│   │   └── service/       # Business logic
│   └── migrations/        # Database migrations
├── assets/                # Brand assets and logos
└── README.md
```

## Prerequisites

| Requirement | Version |
|-------------|---------|
| Node.js | 18+ |
| npm | 9+ |
| Go | 1.23+ |
| GCC | Required for SQLite |

## Getting Started

### Frontend

```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

The frontend will be available at `http://localhost:3000`

### Backend

```bash
cd backend

# Download dependencies
go mod download

# Run the server
make run
# or
go run ./cmd/api
```

The API will be available at `http://localhost:8080`

## Scripts Reference

### Frontend Commands

| Command | Description |
|---------|-------------|
| `npm run dev` | Start development server with HMR |
| `npm run build` | Build for production |
| `npm run preview` | Preview production build locally |
| `npm run lint` | Run OxLint and ESLint |
| `npm run format` | Format code with Prettier |
| `npm run typecheck` | Run TypeScript type checking |

### Backend Commands

| Command | Description |
|---------|-------------|
| `make run` | Run the application |
| `make build` | Build the application binary |
| `make test` | Run test suite |
| `make tidy` | Tidy Go module dependencies |
| `make fmt` | Format Go code |

## API Reference

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/questions` | Retrieve all questionnaire items |
| `POST` | `/api/results` | Submit answers and calculate personality scores |
| `GET` | `/api/results/{id}` | Retrieve a specific result by ID |
| `GET` | `/health` | Health check endpoint |

## Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `DATABASE_PATH` | `./voca.db` | SQLite database file path |
| `ENVIRONMENT` | `development` | Environment mode (`development` / `production`) |

### Database

The backend uses **SQLite** for data persistence. The database file is automatically created when the server starts for the first time. No manual setup required — migrations run automatically on startup.

## Technology Stack

### Frontend
- **[Nuxt 4](https://nuxt.com/)** — Vue 3 meta-framework
- **[TypeScript 5.9](https://www.typescriptlang.org/)** — Type-safe JavaScript
- **[Nuxt UI 4](https://ui.nuxt.com/)** — Component library
- **[Tailwind CSS 4](https://tailwindcss.com/)** — Utility-first CSS
- **[OxLint](https://oxc.rs/)** + **[Prettier](https://prettier.io/)** — Linting and formatting

### Backend
- **[Go 1.23](https://go.dev/)** — High-performance backend
- **Standard library HTTP server** — Minimal dependencies
- **[SQLite](https://www.sqlite.org/)** via go-sqlite3 — Embedded database
- **[OpenAI API](https://openai.com/)** — Career recommendation generation

---

<p align="center">
  <sub>Built with care for the next generation of career explorers</sub>
</p>

## License

MIT
