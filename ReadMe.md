# GoToWeb

An application for discovering GoLand and DevOps with Docker and Kubernetes.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- GoLand: Download and install GoLand from [here](https://www.jetbrains.com/go/download/).
- Docker: Install Docker by following the instructions [here](https://www.docker.com/get-started).
- Kubernetes: [here](https://kubernetes.io/docs/setup/).

## Getting Started

To get started with GoToWeb, follow these steps:

### Prerequisites for Local Development without docker

If you are not using Docker for local development, follow these additional steps:

1. Modify the MongoDB connection URI in the backend code:

   In `backend/main.go`, change the following line from:

   ```go
   clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
   ```
   to:
   ```go
   clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
   ```
   This change allows your backend to connect to a MongoDB instance running on your local machine.

## Getting Started

1. Clone the repository:

    ```shell
    git clone https://github.com/yourusername/GoToWeb.git
    cd GoToWeb
    ```

2. Build and run the backend:

    ```shell
    cd backend
    go build -o main .
    ./main
    ```

3. Build and run the frontend:

    ```shell
    cd ../frontend
    go build -o main .
    ./main
    ```

4. Access the web app:

    - Frontend: [http://localhost:80](http://localhost:80)
    - Backend: [http://localhost:3000](http://localhost:3000)

## Directory Structure

- `backend`: Contains the backend Go code.
- `frontend`: Contains the frontend Go code.
- `static`: Contains static files (HTML, CSS, JS) for the frontend.
- `Dockerfile.backend`: Dockerfile for the backend.
- `Dockerfile.frontend`: Dockerfile for the frontend.
- `Dockerfile.database`: Dockerfile for the database.
- `go.mod` and `go.sum`: Go modules for managing dependencies.
