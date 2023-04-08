# Cocktail Grid

## Setup

You have a choice of two development setups:
- Bare metal (run on your own machine)
- Docker (run in a Docker container) - **recommended**

### Bare metal development environment

#### Requirements

- Go (v1.20)
- Node (v16)
- Yarn (v1.22)

#### Steps

- from the `frontend` directory:
  ```bash
    yarn
    yarn dev
  ```

- from the `backend` directory:
  ```bash
    go run .
  ```

### Docker development environment

#### Requirements

- Docker
- Docker Desktop (optional)

#### Steps

- from the root directory:
  ```bash
    docker compose up
  ```

## Future Development

- replace Reflex with Mutagen?