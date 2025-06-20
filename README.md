

# Todo API

A simple RESTful API for managing tasks, built with Go and the Gin framework. This API allows you to create, read, update, and list tasks.

---

## Features

* **Get all tasks** — `GET /tasks`
* **Get a task by ID** — `GET /tasks/:id`
* **Create a new task** — `POST /tasks`
* **Update an existing task** — `PUT /tasks/:id`

---

## Getting Started

### Prerequisites

* Go 1.18 or newer installed


### Running the API

1. Clone the repository (or create your own file):

```bash
git clone https://github.com/santoridev/todoapi.git
cd todoapi
```

2. Download dependencies
```bash
go mod tidy
```


3. Run the server:

```bash
go run .
```

The server will listen on `http://localhost:8080`.

---

## API Usage

### Get All Tasks

```bash
curl http://localhost:8080/tasks
```

### Get Task by ID

```bash
curl http://localhost:8080/tasks/1
```

### Create a Task

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy groceries", "done": false}'
```

### Update a Task

```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy groceries and fruits", "done": true}'
```

---

## Data Model

```json
{
  "id": 1,
  "title": "Task title",
  "done": false
}
```

* `id`: Unique identifier for the task
* `title`: Description or title of the task
* `done`: Boolean flag indicating if the task is completed

---

## Notes

* The API stores tasks in memory, so all data will be lost when the server restarts.
* Task IDs are assigned incrementally starting from 1.
* No authentication or persistent storage is implemented in this simple demo.

