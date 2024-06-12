# Hello World API

This is a simple "Hello World" API written in Golang.

## Getting Started

### Prerequisites

- Go 1.18 or later
- Docker (optional, for containerization)

### Running Locally

1. Clone the repository:

    ```sh
    git clone https://github.com/ntsedemoorg/hello-world-api.git
    cd hello-world-api
    ```

2. Build and run the application:

    ```sh
    go run main.go
    ```

3. The API will be available at `http://localhost:8080/api/hello`.

### Docker

1. Build the Docker image:

    ```sh
    docker build -t hello-world-api:latest .
    ```

2. Run the Docker container:

    ```sh
    docker run -p 8080:8080 hello-world-api:latest
    ```

3. The API will be available at `http://localhost:8080/api/hello`.

### GitHub Actions

This repository includes a GitHub Actions workflow to build and push the Docker image to Docker Hub. Ensure you have the following GitHub secrets configured:

- `DOCKER_HUB_USERNAME`: Your Docker Hub username.
- `DOCKER_HUB_ACCESS_TOKEN`: Your Docker Hub access token or password.

The workflow will trigger on pushes to the `main` branch.

## API Endpoint

- `GET /api/hello`: Returns a JSON response with a "Hello, World from Golang API!" message.