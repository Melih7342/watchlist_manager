# Watchlist Manager (WLM)

A high-performance microservice built in Go for real-time compliance screening. The service performs fuzzy matching against multiple entity watchlists (Literature, Marvel, Ghibli, Raccoon City) using the Levenshtein distance algorithm. 

To ensure maximum performance, watchlists are indexed in-memory at startup, while a persistent PostgreSQL database is used exclusively to maintain a revision-safe audit trail of all screening requests and their outcomes.

## 🚀 Features

* **In-Memory Rule Engine:** High-speed entity matching loaded from local JSON datasets.
* **Fuzzy Matching:** Custom Levenshtein distance implementation to catch typos, variations, and partial matches across names and aliases.
* **Granular Rule Evaluation:** Independent evaluation pipelines for different watchlists (e.g., `rcRule`, `litRule`).
* **Audit Logging:** Fully automated, persistent logging of every HTTP request and matching decision into a Dockerized PostgreSQL database.
* **Microservice Ready:** Built to integrate seamlessly with other services (e.g., Spring Boot backends) using correlation IDs (`request_id`) and tenancy tracking (`system`).

## 🛠️ Tech Stack

* **Backend:** Go (Standard Library: `net/http`, `encoding/json`, `database/sql`)
* **Database:** PostgreSQL 16 (via `github.com/lib/pq`)
* **Infrastructure:** Docker & Docker Compose

## 📦 Local Setup

### Prerequisites
* [Go](https://golang.org/dl/) (v1.21 or higher recommended)
* [Docker Desktop](https://www.docker.com/products/docker-desktop) (or Docker Engine + Docker Compose)

### 1. Start the Audit Database
The PostgreSQL database runs in a Docker container. To avoid conflicts with other local databases, it is mapped to port **5433** on your host machine.

```bash
# Start the database in the background
docker-compose up -d
```
*Note: The database connection string is pre-configured to `host=localhost port=5433 user=auditor password=supersecret dbname=wlm_audit sslmode=disable`.*

### 2. Run the Go Server
Ensure your watchlist JSON files are correctly formatted and located in the `lists/` directory (`lists/raccoon_city.json`, etc.). 

```bash
# Download dependencies and start the server
go get github.com/lib/pq
go run .
```
The server will start and listen on `http://localhost:9090`.

## 🔌 API Documentation

### `POST /wlm/screen`
Performs a compliance check against all active watchlists.

**Request Payload:**
```json
{
  "request_id": "BOOKING-77429",
  "system": "SPRING_BOOT_HOTEL",
  "first-name": "Albert",
  "last-name": "Wesker",
  "aliases": ["Al", "Captain"],
  "dob": "1960-08-08",
  "nationality": "USA"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:9090/wlm/screen \
-H "Content-Type: application/json" \
-d '{
  "request_id": "BOOKING-77429",
  "system": "SPRING_BOOT_HOTEL",
  "first-name": "Albert",
  "last-name": "Wesker",
  "dob": "1960-08-08"
}'
```

**Response Format:**
The API returns a JSON array containing `MatchResult` objects for each evaluated rule.

```json
[
  {
    "RuleName": "rcRule",
    "IsHit": true,
    "WatchlistID": "001",
    "WatchlistName": "Albert Wesker",
    "Details": {
      "firstName": 100.0,
      "lastName": 100.0,
      "DOB": 100.0
    }
  },
  {
    "RuleName": "litRule",
    "IsHit": false
  }
]
```

## 🗄️ Audit Trail (Database Access)

Every screening request is permanently logged in the `screening_audit` table. You can inspect the logs using any database client (like DBeaver or IntelliJ Database Tool) or directly via the Docker CLI:

```bash
docker exec -it wlm-postgres psql -U auditor -d wlm_audit -c "select * from screening_audit ORDER BY timestamp desc limit 10;"
```
