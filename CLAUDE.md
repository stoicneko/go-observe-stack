# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Docker Compose-based learning project: a Go HTTP backend behind Nginx, with MySQL storage and a full observability stack (Prometheus, Grafana, cAdvisor, Loki, Promtail).

## Architecture

```
Client → Nginx (:8080) → Go app (:5000) → MySQL (:3306)
                              ↑
              Prometheus (:9090) ← cAdvisor (:8081)
              Grafana (:3000) ← Loki (:3100) ← Promtail (Docker logs)
```

- **Go backend** (`main.go`): Single-file HTTP server on `:5000` using `net/http` and `database/sql` with `go-sql-driver/mysql`. Module name: `myapp`, Go 1.18.
- **Nginx** (`nginx.conf`): Reverse proxy on `:80` (host `:8080`) forwarding to `golang:5000`.
- **MySQL 8.0**: Database `myapp`, credentials `root:123456`. Data persisted in `./mysql-data/`.
- **Observability**: Prometheus scrapes itself, Docker daemon (`172.23.0.1:9323`), and cAdvisor. Promtail ships Docker container logs to Loki. Grafana for dashboards (data in `./grafana-data/`).

## Commands

```bash
# Start all services
docker compose up -d

# Rebuild Go app after code changes
docker compose up -d --build golang

# View logs
docker compose logs -f golang
docker compose logs -f nginx

# Stop everything
docker compose down

# Go build/test (local, outside Docker)
go build -o server main.go
go test -race -cover ./...
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/` | Returns "hello from go backend" |
| GET | `/ping-db` | Health check for MySQL connection |
| GET | `/messages` | List all messages from DB |
| POST | `/messages?content=...` | Insert a message into DB |

## Notes

- The Dockerfile uses `GOPROXY=https://goproxy.cn,direct` (Chinese Go module proxy).
- MySQL image is pulled from Alibaba Cloud registry (`registry.cn-hangzhou.aliyuncs.com`).
- `mysql-data/` and `grafana-data/` are Docker volume mounts — do not commit these directories.
