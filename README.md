# go-observe-stack

A learning project for Go backend development, containerization, Kubernetes orchestration, and full-stack observability.

[中文文档](README_CN.md)

## Architecture

```
Client --> Nginx (:8080) --> Go app (:5000) --> MySQL (:3306)
                                 |
                 Prometheus (:9090) <-- cAdvisor (:8081)
                 Grafana (:3000) <-- Loki (:3100) <-- Promtail
```

| Service    | Port  | Description                  |
|------------|-------|------------------------------|
| Nginx      | 8080  | Reverse proxy                |
| Go app     | 5000  | HTTP backend                 |
| MySQL      | 3306  | Data storage                 |
| Prometheus | 9090  | Metrics collection           |
| Grafana    | 3000  | Dashboards & visualization   |
| cAdvisor   | 8081  | Container metrics            |
| Loki       | 3100  | Log aggregation              |
| Promtail   | -     | Log shipping (Docker logs)   |

## Quick Start

```bash
# Start all services
docker compose up -d

# Verify
curl http://localhost:8080/
# => hello from go backend

# Stop
docker compose down
```

## API

| Method | Path                    | Description              |
|--------|-------------------------|--------------------------|
| GET    | `/`                     | Hello message            |
| GET    | `/ping-db`              | MySQL health check       |
| GET    | `/messages`             | List all messages        |
| POST   | `/messages?content=...` | Insert a message         |

## Rebuild After Code Changes

```bash
docker compose up -d --build golang
```

## Observability

- **Metrics**: Prometheus scrapes cAdvisor for container metrics. View at `http://localhost:9090`.
- **Logs**: Promtail collects Docker container logs and ships them to Loki. Query via Grafana.
- **Dashboards**: Grafana at `http://localhost:3000` (default login: admin / admin). Add Prometheus and Loki as data sources.

## Kubernetes Deployment

### Vanilla Manifests

```bash
kubectl apply -f k8s/
```

Includes Deployment + Service for backend, MySQL, and Nginx, plus a ConfigMap for DB connection config. DB password is referenced via a Secret (`backend-secret`) that needs to be created manually.

### Helm Chart

```bash
helm install go-observe-stack ./go-observe-stack
```

The Helm chart (`go-observe-stack/`) packages backend, MySQL, Nginx, Prometheus, and Grafana. Configurable values are in `values.yaml` (image, replicas, port, etc.).

## CI/CD

GitHub Actions pipeline (`.github/workflows/ci.yaml`) runs on every push to `master`:

1. **Test** — `go test ./...`
2. **Build** — Docker image build
3. **Push** — Publish `go-backend:latest` to GitHub Container Registry (`ghcr.io`)

## Testing

```bash
go test ./...
```

Unit tests cover the root handler and input validation for the messages endpoint using `httptest`.

## Project Progress

| Milestone                          | Status |
|------------------------------------|--------|
| Go HTTP backend + MySQL            | Done   |
| Nginx reverse proxy                | Done   |
| Docker Compose orchestration       | Done   |
| Observability (Prometheus, Grafana, Loki, Promtail, cAdvisor) | Done |
| GitHub Actions CI/CD               | Done   |
| Unit tests + CI integration        | Done   |
| Vanilla Kubernetes manifests       | Done   |
| Externalize DB config (ConfigMap/Secret) | Done |
| Helm chart (backend, Prometheus, Grafana) | Done |
| Helm chart — full parameterization (db, nginx) | TODO |
| Integration tests (Testcontainers) | TODO   |
| Multi-stage Dockerfile             | TODO   |
| K8s health probes & resource limits | TODO  |
| Ingress / domain routing           | TODO   |

## License

MIT
