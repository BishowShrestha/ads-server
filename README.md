#  Ad Server Backend

A high-performance, fault-tolerant GoLang backend to **manage**, **track**, and **analyze** video advertisements.

---

## Setup / Run Instructions

### 1. Clone the repository

```bash
git clone https://github.com/BishowShrestha/ad-server.git
cd ad-server
```

### 2. Configure Environment Variables

Create a `.env` file or set environment variables:

| Variable | Description |
|:---------|:------------|
| `DATABASE_URL` | PostgreSQL DB (example: `host=localhost user=postgres password=postgres dbname=ad_server port=5432 sslmode=disable`) |
| `PORT`   | Port to run the server (default: `8080`) |

Example `.env`:

```bash
DB_DSN=host=localhost user=postgres password=postgres dbname=ad_server port=5432 sslmode=disable
PORT=8080
```

### 3. Run the Application

```bash
go mod tidy
go run cmd/server/main.go
```

The server will be available at:  
 `http://localhost:8080`

### 4. Run with Docker

Build and run the Docker container:

```bash
docker build -t ad-server .
docker run -p 8080:8080 ad-server
```

---

##  Concurrency and Multi-Step Processing

###  Efficient Click Handling

- When a click event is received at `/ads/click`, the server:
    - **Pushes the event into a buffered Go channel.**
    - **Immediately responds** to the client without waiting for the database operation.
- A **background Goroutine** listens to the channel:
    - It **reads click events** and **saves** them to the database using GORM.
- This ensures:
    - **Non-blocking** client responses.
    - **High throughput** and efficient handling of spikes.
    - **Resilience** — temporary DB delays won't lose clicks.

###  Benefits
- Super-fast client response time.
- No data loss.
- Easily scalable for viral traffic.

---

##  API Documentation

### 1. Get All Ads

**Endpoint:** `GET /ads`

**Description:**  
Fetch all ads with metadata.

**Response:**

```json
[
  {
    "id": 1,
    "image_url": "https://example.com/ad1.jpg",
    "target_url": "https://targetsite.com/ad1"
  }
]
```

---

###  2. Register a Click

**Endpoint:** `POST /ads/click`

**Description:**  
Register a user click event.

**Request Body:**

```json
{
  "ad_id": 1,
  "ip_address": "192.168.1.1",
  "playback_time": 10.5,
  "timestamp": "2024-04-25T15:04:05Z"
}
```

**Response:**

```json
{
  "message": "Click registered successfully"
}
```

---

###  3. Get Analytics

**Endpoint:** `GET /ads/analytics`

**Description:**  
Get real-time click counts for each ad.

**Response:**

```json
{
  "1": 150,
  "2": 98
}
```
Where:
- Key: `Ad ID`
- Value: `Total Clicks`

---

### 4. Get Hourly Analytics

**Endpoint:** `GET /ads/analytics/hourly`

**Description:**  
Get real-time click counts for each ad.

**Response:**

```json
{
  "hour": "2025-04-28T20",
  "count": 98
}
```
Where:
- Hour: `Hourly Time`
- Click: `Total Clicks`

---


### 5. Prometheus Metrics

**Endpoint:** `GET /metrics`

**Description:**  
Expose metrics for Prometheus scraping, including:

- API request count
- Request durations
- Error rates
- Business metrics (clicks, ads served)

---

##  Example API Flow

1. User fetches ads → `GET /ads`
2. User clicks an ad → `POST /ads/click`
3. Admin fetches performance → `GET /ads/analytics`
4. Admin fetches hourly performance → `GET /ads/analytics/hourly`

---

##  Future Enhancements

- Batch database writes for even higher performance.
- Add Redis buffer for peak loads.
- Implement circuit breakers on DB operations.

---

##  Authors

- **Bishow Shrestha**
- Built with: **Go**, **GORM**, **PostgreSQL**, **Docker**, **Prometheus**

---


