#!/usr/bin/env bash
set -euo pipefail

mkdir -p .test-results
go test ./... -coverprofile=.test-results/coverage.out
