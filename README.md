# 🔍 Goseek 🔍

High-performance product search API built with Go — lightweight, scalable, and blazing fast.

## 📑 Table of Contents

- [✅ Prerequisites](#-prerequisites)
- [🚀 Getting Started](#-getting-started)
  - [Setup](#setup)
  - [Development (Local)](#development-local)
  - [Staging (AWS)](#staging-aws)
  - [Production (AWS)](#production-aws)
  - [Cleanup](#cleanup)
- [📁 Project Structure](#-project-structure)

## ✅ Prerequisites

- **Docker Desktop** - For containerized development
- **Go 1.25.1+** - For local IDE support
- **AWS CLI** - For AWS deployments, configured with `aws configure`
- **Terraform 1.0+** - For infrastructure as code, run `terraform init` in `terraform/stage` and `terraform/prod`
- **VS Code or GoLand** (Optional) - For IDE features with autocomplete

## 🚀 Getting Started

### Setup

Downloads Go dependencies for IDE autocomplete and navigation features. Run this once after cloning the repository.

```bash
make setup
```

### Development (Local)

#### Deploy

Builds Docker image with Swagger enabled, starts container with Air hot-reload, and mounts source code as volume. When you edit any `.go` file and save, Air automatically detects changes, regenerates Swagger docs, recompiles the binary, and restarts the application (typically 2-5 seconds).

```bash
make deploy-dev
```

To test the API, open `http://localhost:8080/swagger/index.html` or use `cURL`:

```bash
# Search products by query
curl 'http://localhost:8080/products/search?q=a'

# Get server health
curl 'http://localhost:8080/health'
```

#### Logs

Shows all container logs with real-time streaming. `Ctrl+C` to exit, container keeps running.

```bash
make log-dev
```

#### Shell

Opens interactive shell inside the running container for inspection and debugging. Type `exit` to close.

```bash
make shell-dev
```

#### Stop

Stops and removes the development container. Keeps Docker images for faster restart.

```bash
make stop-dev
```

#### Destroy

Removes containers, volumes, and images. Frees disk space.

```bash
make destroy-dev
```

### Staging (AWS)

#### Deploy

First deployment creates ECR repository and infrastructure. Subsequent deployments build new image, push to ECR, and update ECS service. Waits for service stability and displays public IP.

```bash
make deploy-stage
```

To test the API, open `http://<STAGING_IP>:8080/swagger/index.html` or use `cURL`:

```bash
# Search products by query
curl 'http://<STAGING_IP>:8080/products/search?q=a'

# Get server health
curl 'http://<STAGING_IP>:8080/health'
```

#### Logs

Streams CloudWatch logs in real-time. `Ctrl+C` to exit, service keeps running.

```bash
make log-stage
```

#### Shell

Opens interactive shell in the running ECS task on AWS. Type `exit` to close.

```bash
make shell-stage
```

#### Stop

Scales service to 0 tasks. Infrastructure remains, no compute costs.

```bash
make stop-stage
```

#### Start

Scales service back to 1 task.

```bash
make start-stage
```

#### Destroy

Deletes all AWS staging resources and removes local Docker images. Requires typing `yes` to confirm.

```bash
make destroy-stage
```

### Production (AWS)

#### Deploy

Swagger disabled for security. Separate ECR repository. Waits for service stability and displays public IP.

```bash
make deploy-prod
```

To test the API, use `cURL` (Swagger is disabled in production):

```bash
# Search products by query
curl 'http://<PRODUCTION_IP>:8080/products/search?q=a'

# Get server health
curl 'http://<PRODUCTION_IP>:8080/health'
```

#### Logs

Streams CloudWatch logs from production tasks. `Ctrl+C` to exit, service keeps running.

```bash
make log-prod
```

#### Shell

Opens interactive shell in one of the running production tasks. Type `exit` to close.

```bash
make shell-prod
```

#### Stop

Scales service to 0 tasks. Infrastructure remains.

```bash
make stop-prod
```

#### Start

Scales service back to 1 task.

```bash
make start-prod
```

#### Destroy

Deletes all AWS production resources and removes local Docker images. Requires typing `yes` to confirm.

```bash
make destroy-prod
```

### Cleanup

Removes local build artifacts (dist/, tmp/, docs/). Does not affect Docker containers or AWS resources.

```bash
make clean
```

## 📁 Project Structure

```
Goseek/
├── .vscode/
├── cmd/
│   └── api/
│       ├── main.go
│       ├── swagger_prod.go
│       └── swagger.go
├── data/
│   └── products.json
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── handlers/
│   │   ├── products_handler.go
│   │   └── root_handler.go
│   ├── models/
│   │   ├── http_error.go
│   │   ├── http_response.go
│   │   └── product.go
│   ├── repository/
│   │   └── products_repository.go
│   ├── router/
│   │   └── router.go
│   ├── services/
│   │   └── products_service.go
│   └── shared/
│       └── shared_error.go
├── terraform/
│   ├── modules/
│   │   ├── alb/
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   └── variables.tf
│   │   ├── ecr/
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   └── variables.tf
│   │   ├── ecs/
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   └── variables.tf
│   │   ├── iam/
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   └── variables.tf
│   │   ├── networking/
│   │   │   ├── main.tf
│   │   │   ├── outputs.tf
│   │   │   └── variables.tf
│   │   └── security/
│   │       ├── main.tf
│   │       ├── outputs.tf
│   │       └── variables.tf
│   ├── prod/
│   │   ├── .terraform.lock.hcl
│   │   ├── main.tf
│   │   ├── outputs.tf
│   │   ├── terraform.tfvars
│   │   └── variables.tf
│   └── stage/
│       ├── .terraform.lock.hcl
│       ├── main.tf
│       ├── outputs.tf
│       ├── terraform.tfvars
│       ├── terraform.tfvars.example
│       └── variables.tf
├── tmp/
├── .air.toml
├── .gitignore
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```
