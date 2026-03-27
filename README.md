# go-observe-stack

A Docker Compose project: Go HTTP backend behind Nginx, with MySQL storage and a full observability stack.

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

## License

MIT
