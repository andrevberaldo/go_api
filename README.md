## Quick sandbox CRUD Rest API for personal learning using Go and Gin Gonic

#### This project aimed to serve as a sandbox for me to practice Go's syntax and explore its unique error-handling concept. I focused on understanding error flow and building the /products route, implementing the main HTTP verbs while deepening my knowledge of the Gin Gonic framework. Some concepts are only simulated and not fully implemented (such as JWT authentication). Another important point to highlight is that configuration values are exposed, which would not be acceptable in a real application.

![alt text](image-1.png)

## How to Run

 - Docker Engine and Docker compose is needed.
 - Run `docker compose up` in the terminal
 - Once the container is running
 - The PGAdmin will start at localhost:8080, you can connect with postgres DB and run the following script to create products table.
 ```sql
 CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price NUMERIC (10, 2)
 )
 ```
 - Run `go run ./cmd/main.go` 
 - To use the **protected routes** on /api path, the request headers must have the fake JWT at `Authorization: fake_token_JWT`

 ## Next steps
 - Unit test
 - Lint 