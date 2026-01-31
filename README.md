# Voca - Personality Test

A web application for administering the IPIP-50 (International Personality Item Pool) personality assessment based on the Big Five model.

## Features

- Step-by-step questionnaire with progress indicator
- Big Five personality trait assessment (Extraversion, Agreeableness, Conscientiousness, Emotional Stability, Openness)
- Visual results with trait scores and interpretations
- Responsive design for desktop and mobile
- Multi-language support (10 languages)

## Scientific Foundation

This assessment uses the **IPIP-50**, a validated 50-item scale from the International Personality Item Pool (IPIP). IPIP is a public-domain collection of personality assessment items developed by leading psychology researchers. Unlike proprietary tests, IPIP instruments are freely available for research and practical use.

Learn more at [ipip.ori.org](https://ipip.ori.org/)

## Project Structure

```
voca/
├── frontend/          # Nuxt 3 application (TypeScript, Nuxt UI)
│   ├── app/
│   │   ├── components/
│   │   ├── composables/
│   │   └── pages/
│   └── ...
├── backend/           # Go API server
│   ├── cmd/api/       # Application entry point
│   ├── internal/
│   │   ├── config/    # Configuration
│   │   ├── domain/    # Domain models
│   │   ├── handler/   # HTTP handlers
│   │   ├── repository/# Data access
│   │   └── service/   # Business logic
│   └── migrations/    # Database migrations
└── README.md
```

## Prerequisites

- Node.js 18+ and npm
- Go 1.21+
- GCC (required for SQLite compilation)

## Getting Started

### Frontend

```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

The frontend will be available at `http://localhost:3000`.

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

The API will be available at `http://localhost:8080`.

## Frontend Scripts

| Command | Description |
|---------|-------------|
| `npm run dev` | Start development server |
| `npm run build` | Build for production |
| `npm run preview` | Preview production build |
| `npm run lint` | Run OxLint and ESLint |
| `npm run format` | Format code with Prettier |
| `npm run typecheck` | Run TypeScript type checking |

## Backend Makefile Commands

| Command | Description |
|---------|-------------|
| `make run` | Run the application |
| `make build` | Build the application |
| `make test` | Run tests |
| `make tidy` | Tidy dependencies |
| `make fmt` | Format code |

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/questions` | Get all questionnaire items |
| POST | `/api/results` | Submit answers and calculate results |
| GET | `/api/results/{id}` | Get a specific result by ID |
| GET | `/health` | Health check endpoint |

## Environment Variables

### Backend

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `DATABASE_PATH` | `./voca.db` | SQLite database file path |
| `ENVIRONMENT` | `development` | Environment (development/production) |

## Database

The backend uses SQLite for data persistence. The database file is automatically created when the server starts for the first time. By default, it's stored at `./voca.db` in the backend directory.

No manual setup is required - migrations run automatically on startup.

## Technology Stack

### Frontend
- Nuxt 3 (Vue 3)
- TypeScript
- Nuxt UI (Tailwind CSS v4)
- OxLint + Prettier

### Backend
- Go 1.21+
- Standard library HTTP server
- SQLite (via go-sqlite3)

## License

MIT
