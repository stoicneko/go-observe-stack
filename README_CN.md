# go-observe-stack

基于 Docker Compose 的学习项目：Go HTTP 后端 + Nginx 反向代理 + MySQL 存储 + 全套可观测性组件。

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

## 许可证

MIT
