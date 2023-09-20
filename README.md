# Blog Infrastructure with REST API, AWS SQS, Redis Caching, and PostgreSQL

This repository contains the backend infrastructure and code for setting up a blog system with a RESTful API. The system uses AWS SQS (Simple Queue Service) for handling asynchronous tasks, Redis for caching frequently accessed data, and PostgreSQL as the persistent database.

## Prerequisites

Before you begin, make sure you have the following prerequisites installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://go.dev/)

## Getting Started

Follow these steps to set up and run the blog infrastructure:

1. **Clone this repository to your local machine:**

   ```bash
   git clone https://github.com/joshjms/monokrom.git
   cd monokrom
   ```

2. **Create a copy of `.env.example`**

3. **Build and start the Docker containers using `docker-compose`**

    ```bash
    docker-compose up -d
    ```

4. **Run the Go application**

    ```bash
    go run .
    ```

Your application should be running at `localhost:3000`.

## API Documentation

- `POST` - `/post`
    Creates a new post.

- `GET` - `/post/:slug`
    Gets a post.