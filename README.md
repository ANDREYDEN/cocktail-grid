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

Optionally you can run the containers separately:
- from the root directory:
  ```bash
    docker compose up frontend
    docker compose up backend
  ```

## Deployment

In a production environment the `frontend` and `backend` containers can be started with:
```
docker-compose up -f docker-compose.prod.yml -d
```

Sometimes a fresh build might be required:
```
docker-compose up -f docker-compose.prod.yml -d --build
```

In case the machine runs out of space, dangling docker resources can be cleaned up:
```
docker system prune
```

## References

- [Go containarization](https://karol-filipczuk.medium.com/development-environment-for-web-app-with-containers-docker-vue-js-go-part-1-b0f4d663bd64)
- [Vue containarization](https://karol-filipczuk.medium.com/development-environment-for-web-app-with-containers-docker-vue-js-go-part-2-25f059d88cb8)

## Future Development

- replace Reflex with Mutagen?