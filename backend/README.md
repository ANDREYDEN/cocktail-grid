# Backend

## Prerequisites
- Go 1.20+
  - Add GOPATH to PATH: `export PATH=$(go env GOPATH)/bin:$PATH` 
- [`swagger-codegen`](https://github.com/swagger-api/swagger-codegen/tree/master): `brew install swagger-codegen`
- [`swag`](https://github.com/swaggo/swag): `go install github.com/swaggo/swag/cmd/swag@latest`

## Development

To get the backend running execute the following commands:
1. Set up `.env`: `cp .env.example .env`
2. Install dependencies: `go get .`
3. Run the server: `make run`

## Useful commands

1. Run the server
    
    ```bash
    go run .
    ```

2. Run the seeders

    ```bash
    go run . --seed
    ```

3. Generate swagger docs

    ```bash
    swag init
    ```

4. Format swagger docs
    
    ```bash
    swag fmt 
    ```