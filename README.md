---
title: kube-app-manager
description: Service that exposes Kubernetes WebEndpoint custom resources through a simple HTTP API.
---

[![Tests](https://img.shields.io/github/actions/workflow/status/xkovacikm2/kube-app-manager/ci.yml?branch=main&label=tests)](https://github.com/xkovacikm2/kube-app-manager/actions/workflows/ci.yml)
[![Build](https://img.shields.io/github/actions/workflow/status/xkovacikm2/kube-app-manager/ci.yml?branch=main&label=build)](https://github.com/xkovacikm2/kube-app-manager/actions/workflows/ci.yml)

## Overview
kube-app-manager is a small Go service that reads WebEndpoint custom resources
(group apps.kovko.top, version v1alpha1, resource webendpoints) from
Kubernetes and returns them as JSON.

The service exposes one endpoint:
- GET /apps

## What It Is For
This project provides a lightweight API layer for app catalogs stored in
Kubernetes custom resources. It is useful when you want a simple HTTP interface
for UI or automation tools without embedding Kubernetes client logic in each
consumer.

## Prerequisites
- Go 1.23+
- Access to a Kubernetes cluster that contains the WebEndpoint CRD
- A valid Kubernetes config:
  - In-cluster config when running in Kubernetes, or
  - local kubeconfig via default location or KUBECONFIG
- Optional: Task CLI for convenience commands
- Optional: Docker for container build and run

## Run Locally
### Option 1: Go directly
```bash
go run ./cmd/kube-app-manager
```

### Option 2: Task command
```bash
task start
```

The service listens on port 8080 by default. Override with PORT.

Quick check:
```bash
curl http://localhost:8080/apps
```

## Run Tests
```bash
go test ./...
```

Or with task:
```bash
task test
```

## Build Binary
```bash
go build -o .bin/kube-app-manager ./cmd/kube-app-manager
```

Or with task:
```bash
task build
```

## Build and Run Docker Image
Build from repository root because the Dockerfile is under build/docker and
uses project-root build context.

```bash
docker build -f build/docker/Dockerfile -t kube-app-manager:local .
```

Run:
```bash
docker run --rm -p 8080:8080 \
  -v "$HOME/.kube:/home/nonroot/.kube:ro" \
  -e KUBECONFIG=/home/nonroot/.kube/config \
  kube-app-manager:local
```

## Deploy with Helm
Helm chart is available in charts/kube-app-manager.

```bash
helm upgrade --install kube-app-manager ./charts/kube-app-manager
```

## CI
GitHub Actions workflow is in .github/workflows/ci.yml and enforces:
- Tests run first
- Docker build runs only after tests pass
- Image push is enabled on push to main
