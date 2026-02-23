# Ride Simulator

A distributed ride-sharing simulator built with microservices architecture, featuring a **User Service (Go)** and **Ride Service (Laravel/PHP)**, backed by PostgreSQL and Redis.

---
## API Documentation

https://docs.google.com/document/d/18509p3yzW9mcHLRz36KLpBOvBoqTIZnkPa8XGa98mck/edit?usp=sharing

---

## Project Architecture

```
RideSimulator/
├── docker-compose.yaml          # Docker Compose setup (v2)
├── kind-config.yaml             # KIND cluster configuration
├── setup-kind.sh                # KIND cluster setup script
├── port-forward.sh              # Port forwarding script
├── init-db.sql                  # Database initialization script
├── README.md                    # This file
├── User/                        # Go User Service
│   ├── Dockerfile
│   ├── .env
│   ├── main.go
│   ├── go.mod
│   └── ...
├── Ride/                       # Laravel Ride Service
│   ├── Dockerfile
│   ├── .env.example
│   └── ...
└── k8s/                        # Kubernetes manifests
    ├── postgres.yaml
    ├── redis.yaml
    ├── user-service.yaml
    └── rider-service.yaml
```

---

## Services Overview

### 1. User Service (Go)

- **Container Name**: `user-app-v2`
- **Port**: 8085 (Docker), 8080 (KIND)
- **Purpose**: User registration and authentication for riders and drivers
- **Features**:
  - Rider registration
  - Driver registration
  - User status management (online/offline)
  - Basic authentication protection

### 2. Ride Service (Laravel/PHP)

- **Container Name**: `ride-app-v2`
- **Port**: 8005 (Docker), 8000 (KIND)
- **Purpose**: Ride management and tracking
- **Features**:
  - Ride request creation
  - Ride status tracking (started/ended)
  - Rider and driver ride history

### 3. PostgreSQL Database

- **Container Name**: `pg-container-v2`
- **Port**: 5433 (Docker), 5432 (KIND)
- **Databases**: `userdb`, `riderdb`

### 4. Redis Cache

- **Container Name**: `redis-stack-v2`
- **Port**: 6380 (Docker), 6379 (KIND)

---

## Quick Start

### Option 1: Docker Compose (Recommended for Development)

```bash
# Build and start all services
docker-compose up -d --build

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Option 2: Kubernetes (KIND) - Recommended for Testing

```bash
# Run the setup script
./setup-kind.sh

# Start port forwarding in a new terminal
./port-forward.sh
```

---

## Kubernetes Setup (KIND)

### Prerequisites

- Docker
- Kind
- kubectl

### Automated Setup

```bash
# Run the setup script - this will:
# 1. Create a KIND cluster
# 2. Build and push Docker images to Docker Hub
# 3. Deploy all services
./setup-kind.sh
```

### Manual Setup (If Needed)

1. **Create Cluster**:

```bash
kind create cluster --config kind-config.yaml --wait 3m
```

2. **Build & Load Images**:

```bash
docker build -t ridesimulator-user-app-v2:latest ./User
docker build -t ridesimulator-ride-app-v2:latest ./Ride

kind load docker-image ridesimulator-user-app-v2:latest --name ride-cluster
kind load docker-image ridesimulator-ride-app-v2:latest --name ride-cluster
```

3. **Deploy Services**:

```bash
kubectl apply -f k8s/postgres.yaml
kubectl apply -f k8s/redis.yaml
kubectl apply -f k8s/user-service.yaml
kubectl apply -f k8s/rider-service.yaml
```

4. **Start Port Forwarding**:

```bash
./port-forward.sh
```

---

## Accessing Services

### Docker Compose

| Service       | URL                          |
|---------------|------------------------------|
| User Service  | http://localhost:8085        |
| Ride Service  | http://localhost:8005        |
| PostgreSQL    | localhost:5433               |
| Redis         | localhost:6380              |

### Kubernetes (KIND)

| Service       | URL                          |
|---------------|------------------------------|
| User Service  | http://localhost:8080        |
| Ride Service  | http://localhost:8000        |

---

## API Documentation

### User Service API

All endpoints are protected with Basic Authentication.

**Base URL**: `http://localhost:8085` (Docker) or `http://localhost:8080` (KIND)

**Authentication**:
- Username: `admin`
- Password: `admin123`

#### Register a Rider

```bash
curl -X POST "http://localhost:8085/api/v1/riders" \
-H "Content-Type: application/json" \
-d '{"phone":"01712345669","email":"johny@gmail.com","name":"Johny Depp"}'
```

#### Register a Driver

```bash
curl -X POST "http://localhost:8085/api/v1/drivers" \
-H "Content-Type: application/json" \
-d '{"phone":"01712345670","email":"driver@example.com","name":"Johnyy Depp"}'
```

#### Update Driver Status

```bash
curl -X PUT "http://localhost:8085/api/v1/drivers/{id}/status" \
-H "Content-Type: application/json" \
-d '{"status":"online"}'
```

#### Get User Status

```bash
curl -X GET "http://localhost:8085/api/v1/users/{id}/status" \
-H "Authorization: Basic YWRtaW46YWRtaW4xMjM="
```

---

### Ride Service API

**Base URL**: `http://localhost:8005` (Docker) or `http://localhost:8000` (KIND)

#### Request a Ride

```bash
curl -X POST "http://localhost:8005/api/v1/rides" \
-H "Content-Type: application/json" \
-d '{"rider_id":1}'
```

#### End a Ride

```bash
curl -X PUT "http://localhost:8005/api/v1/rides" \
-H "Content-Type: application/json" \
-d '{"id":1}'
```

#### Get Rider's Rides

```bash
curl -X GET "http://localhost:8005/api/v1/rides?rider_id={id}"
```

#### Get Driver's Rides

```bash
curl -X GET "http://localhost:8005/api/v1/rides?driver_id={id}"
```

---

## Database Schema

### userdb Database

**users Table**

| Field          | Type                     | Description                          |
|----------------|--------------------------|--------------------------------------|
| id             | SERIAL PRIMARY KEY      | Unique auto-generated identifier    |
| name           | VARCHAR(255)             | User's full name                     |
| mobile_no      | VARCHAR(20) UNIQUE      | Must be unique                       |
| email          | VARCHAR(255) UNIQUE     | Must be unique                       |
| user_type      | ENUM ('rider', 'driver')| Type of user                        |
| cur_status     | ENUM ('online', 'offline')| Current status (default: offline) |
| password       | VARCHAR(255)            | Hashed password (optional)           |
| created_at     | TIMESTAMP               | Record creation timestamp            |
| updated_at     | TIMESTAMP               | Record last update timestamp         |

### riderdb Database

**rides Table**

| Field      | Type                     | Description                                     |
|------------|--------------------------|------------------------------------------------|
| id         | SERIAL PRIMARY KEY      | Unique auto-generated ride ID                   |
| rider_id   | INTEGER                 | ID of rider requesting the ride                 |
| driver_id  | INTEGER                 | ID of driver assigned                           |
| status     | ENUM ('started', 'ended')| Ride status                                    |
| started_at | TIMESTAMP               | Timestamp when ride began                       |
| ended_at   | TIMESTAMP               | Timestamp when ride ended (optional if ongoing) |
| created_at | TIMESTAMP               | Record creation timestamp                       |
| updated_at | TIMESTAMP               | Record last updated timestamp                   |

---

## Environment Configuration

### User Service (`.env`)

```bash
DB_HOST=pg-container-v2
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=userdb
SERVER_PORT=8080
BASIC_AUTH_USER=admin
BASIC_AUTH_PASSWORD=admin123
REDIS_HOST=redis-stack-v2
REDIS_PORT=6379
```

### Kubernetes Environment Variables

**User Service**:
```bash
DB_HOST=postgres-service
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=db01
DB_NAME=UserDB
SERVER_PORT=8080
BASIC_AUTH_USER=admin
BASIC_AUTH_PASSWORD=admin123
REDIS_HOST=redis-service
REDIS_PORT=6379
```

**Ride Service**:
```bash
APP_ENV=production
APP_DEBUG=false
APP_URL=http://localhost:8000
DB_CONNECTION=pgsql
DB_HOST=postgres-service
DB_PORT=5432
DB_DATABASE=riderdb
DB_USERNAME=postgres
DB_PASSWORD=db01
EXTERNAL_API_URL=http://user-service:8080
EXTERNAL_API_USERNAME=admin
EXTERNAL_API_PASSWORD=admin123
REDIS_HOST=redis-service
REDIS_PORT=6379
CACHE_STORE=database
SESSION_DRIVER=database
QUEUE_CONNECTION=database
```

---

## Helper Scripts

### `setup-kind.sh`

Automates the entire KIND cluster setup:
- Creates/Recreates the cluster
- Builds Docker images
- Pushes images to Docker Hub
- Deploys all Kubernetes resources
- Waits for services to be ready

### `port-forward.sh`

Establishes port forwarding to access services:
- User Service: localhost:8080 → user-service:8080
- Ride Service: localhost:8000 → ride-service:8000

---

## Troubleshooting

### Common Issues

1. **Port Conflicts**:
   - If ports 8080, 8000, 5432, or 6379 are in use, stop other services or modify the port mappings.

2. **Database Connection Issues**:
   - Ensure PostgreSQL is running and ready before starting the application services.
   - Check that the database name, username, and password match across configurations.

3. **Image Pull Issues**:
   - Ensure you're logged in to Docker Hub (`docker login`) before running `setup-kind.sh`.
   - For local testing, use `kind load docker-image` instead of pushing to Docker Hub.

4. **KIND Cluster Issues**:
   - Delete and recreate the cluster: `kind delete cluster --name ride-cluster`
   - Check cluster status: `kind get clusters`

### Debugging Commands

```bash
# View all pods
kubectl get pods

# View pod logs
kubectl logs deployment/user-service
kubectl logs deployment/ride-service
kubectl logs deployment/postgres

# Describe resources
kubectl describe pod <pod-name>
kubectl describe service <service-name>

# View resource usage
kubectl top pods

# Access pod shell
kubectl exec -it <pod-name> -- /bin/sh
```

---

## Security Notes

- **Change default passwords** in production environments
- Use proper SSL certificates for HTTPS
- Implement proper authentication and authorization
- Regularly update Docker images and dependencies
- The Basic Auth credentials (`admin`/`admin123`) should be changed in production

---

## Tech Stack

- **User Service**: Go, PostgreSQL, Redis
- **Ride Service**: Laravel (PHP 8.2+), PostgreSQL, Redis
- **Container Orchestration**: Docker Compose, Kubernetes (KIND)
- **Database**: PostgreSQL
- **Cache**: Redis

---

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes
4. Test with Docker Compose or KIND
5. Submit a pull request

---

## License

This project is open-source and available for learning and development purposes.
