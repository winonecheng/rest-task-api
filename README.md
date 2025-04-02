# REST Task API

A simple RESTful API for task management built with Go and Gorilla Mux.

## Overview

This project provides a lightweight HTTP API for managing tasks with the following features:
- Create, read, update, and delete tasks
- Task status tracking (incomplete/complete)
- Concurrent request handling with proper synchronization
- Containerized deployment

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/tasks` | Retrieve all tasks |
| POST | `/tasks` | Create a new task |
| PUT | `/tasks/{id}` | Update an existing task |
| DELETE | `/tasks/{id}` | Delete a task |

## Task Model

```json
{
  "id": 1,
  "name": "Example task",
  "status": 0
}
```

- `id`: Unique identifier (auto-assigned)
- `name`: Task description
- `status`: Task status (0=incomplete, 1=complete)

## Running the API

### Local Development

```bash
# Clone the repository
git clone https://github.com/yourusername/rest-task-api.git
cd rest-task-api

# Run the API
go run main.go
```

### Using Docker

```bash
# Build the Docker image
docker build -t rest-task-api .

# Run the container
docker run -d -p 8080:8080 rest-task-api
```

## API Usage Examples

You can interact with the API using curl commands:

### List all tasks
```bash
curl -X GET http://localhost:8080/tasks
```

### Create a new task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"name": "Complete project", "status": 0}'
```

### Update a task
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Complete project", "status": 1}'
```

### Delete a task
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## Testing

Run the included tests with:

```bash
go test -v
```

## Implementation Details

- In-memory task storage with thread-safe access
- Statically compiled Go binary for minimal container size
- Multi-stage Docker build for optimized image size
