# go-observe-stack

Go 后端开发、容器化、Kubernetes 编排与全栈可观测性学习项目。

[English](README.md)

## 架构

```
Client --> Nginx (:8080) --> Go app (:5000) --> MySQL (:3306)
                                 |
                 Prometheus (:9090) <-- cAdvisor (:8081)
                 Grafana (:3000) <-- Loki (:3100) <-- Promtail
```

| 服务       | 端口  | 说明                     |
|------------|-------|--------------------------|
| Nginx      | 8080  | 反向代理                 |
| Go app     | 5000  | HTTP 后端服务            |
| MySQL      | 3306  | 数据存储                 |
| Prometheus | 9090  | 指标采集                 |
| Grafana    | 3000  | 仪表盘与可视化           |
| cAdvisor   | 8081  | 容器指标                 |
| Loki       | 3100  | 日志聚合                 |
| Promtail   | -     | 日志采集（Docker 日志）  |

## 快速开始

```bash
# 启动所有服务
docker compose up -d

# 验证
curl http://localhost:8080/
# => hello from go backend

# 停止
docker compose down
```

## API 接口

| 方法 | 路径                    | 说明               |
|------|-------------------------|--------------------|
| GET  | `/`                     | 返回欢迎信息       |
| GET  | `/ping-db`              | MySQL 健康检查     |
| GET  | `/messages`             | 查询所有消息       |
| POST | `/messages?content=...` | 插入一条消息       |

## 修改代码后重新构建

```bash
docker compose up -d --build golang
```

## 可观测性

- **指标监控**：Prometheus 通过 cAdvisor 采集容器指标，访问 `http://localhost:9090` 查看。
- **日志系统**：Promtail 采集 Docker 容器日志并发送到 Loki，通过 Grafana 查询。
- **可视化面板**：Grafana 地址 `http://localhost:3000`（默认账号：admin / admin），添加 Prometheus 和 Loki 作为数据源即可。

## Kubernetes 部署

### 原生 Manifest

```bash
kubectl apply -f k8s/
```

包含 backend、MySQL、Nginx 的 Deployment + Service，以及 DB 连接配置的 ConfigMap。数据库密码通过 Secret（`backend-secret`）引用，需手动创建。

### Helm Chart

```bash
helm install go-observe-stack ./go-observe-stack
```

Helm Chart（`go-observe-stack/`）打包了 backend、MySQL、Nginx、Prometheus 和 Grafana，可通过 `values.yaml` 配置镜像、副本数、端口等参数。

## CI/CD

GitHub Actions 流水线（`.github/workflows/ci.yaml`）在每次推送到 `master` 时触发：

1. **测试** — `go test ./...`
2. **构建** — Docker 镜像构建
3. **推送** — 发布 `go-backend:latest` 到 GitHub Container Registry（`ghcr.io`）

## 测试

```bash
go test ./...
```

使用 `httptest` 编写单元测试，覆盖根路由处理和消息接口的输入校验。

## 项目进度

| 里程碑                                  | 状态   |
|-----------------------------------------|--------|
| Go HTTP 后端 + MySQL                    | 已完成 |
| Nginx 反向代理                          | 已完成 |
| Docker Compose 编排                     | 已完成 |
| 可观测性（Prometheus、Grafana、Loki、Promtail、cAdvisor） | 已完成 |
| GitHub Actions CI/CD                    | 已完成 |
| 单元测试 + CI 集成                      | 已完成 |
| 原生 Kubernetes Manifest                | 已完成 |
| 外部化 DB 配置（ConfigMap/Secret）       | 已完成 |
| Helm Chart（backend、Prometheus、Grafana）| 已完成 |
| Helm Chart — 完整参数化（db、nginx）     | 待完成 |
| 集成测试（Testcontainers）               | 待完成 |
| 多阶段 Dockerfile                       | 待完成 |
| K8s 健康探针与资源限制                   | 待完成 |
| Ingress / 域名路由                      | 待完成 |

## 许可证

MIT
