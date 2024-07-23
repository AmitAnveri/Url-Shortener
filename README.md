# URL Shortener (Cloud-Native, Scalable)

This is a cloud-native, scalable URL shortener built with Golang (Gin Framework), PostgreSQL, and Redis, containerized using Docker. It generates unique Base62 short URLs and supports caching with Redis for high performance.

---

## Features
- Scalable - Built with microservices architecture
- Fast - Uses Redis caching for quick lookups
- Efficient - Uses Base62 encoding for unique short URLs
- Cloud-Native - Deployed with Docker and Kubernetes
- Database-Powered - Uses PostgreSQL for persistence

---

## Tech Stack
- Golang (Gin Framework)
- PostgreSQL (Database)
- Redis (Caching)
- Docker and Docker Compose
- Kubernetes
- Base62 Encoding
- GitHub Actions (for CI/CD)

---

## Project Structure
```
/url-shortener
│── services/
│   ├── shortener/
│   │   ├── main.go          # Entry point of the app
│   │   ├── config.go        # Environment variables
│   │   ├── handlers.go      # API request handlers
│   │   ├── storage.go       # Database interactions
│   │   ├── cache.go         # Redis caching functions
│   │   ├── utils.go         # URL encoding and helpers
│   │   ├── models.go        # Data models
│   │   ├── migrate.go       # Database migrations
│   │   ├── Dockerfile       # Containerization
│   │   ├── go.mod & go.sum  # Dependencies
│── migrations/              # Database schema migrations
│── deployments/             # Kubernetes deployment files
│── docker-compose.yml       # Local deployment setup
│── README.md                # Documentation
```

## Getting Started
### Clone the Repository
```sh
git clone https://github.com/YOUR_GITHUB_USERNAME/url-shortener.git
cd url-shortener
```

### Set Up Environment Variables
Create a `.env` file and add:
```ini
DATABASE_URL=postgres://user:password@urlshortener-db:5432/urlshortener?sslmode=disable
REDIS_ADDR=urlshortener-cache:6379
PORT=8080
```

---

## Running with Docker
### Build and Run Containers
```sh
docker-compose up --build -d
```
### Check Running Containers
```sh
docker ps
```

---

## Running the API
If running inside Docker:
```sh
docker-compose up -d
```

---

## API Endpoints
| Method  | Endpoint       | Description |
|---------|---------------|-------------|
| `POST`  | `/shorten`    | Shortens a URL |
| `GET`   | `/:shortUrl`  | Redirects to the original URL |

### Example Usage
#### Shorten a URL
```sh
curl -X POST "http://localhost:8080/shorten" -H "Content-Type: application/json" -d '{"url":"https://example.com"}'
```
#### Redirect to the Original URL
```sh
curl -v "http://localhost:8080/aB3dE9"
```

## License
This project is MIT Licensed. Feel free to use it for your own projects.
