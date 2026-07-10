#!/usr/bin/env bash
set -euo pipefail

docker build -f build/docker/Dockerfile -t kube-app-manager .
