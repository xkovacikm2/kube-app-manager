---
description: 'Best practices for creating optimized, secure, and efficient Docker images and managing containers'
applyTo: '**/Dockerfile,**/Dockerfile.*,**/*.dockerfile,**/docker-compose*.yml,**/docker-compose*.yaml,**/compose*.yml,**/compose*.yaml'
---
# Docker Instructions
Guide developers in building efficient, secure, and maintainable Docker images with emphasis on optimization, security, and reproducibility.

## Core Principles

### Portability
- Design environment-agnostic containers
- Use environment variables for configuration
- Ensure all dependencies are self-contained

### Isolation
- Run single process per container
- Use container networking, not host networking
- Implement resource limits

### Efficiency
- Minimize image size to reduce build time, storage, and attack surface
- Use multi-stage builds and minimal base images
- Regularly optimize and review image size

## Dockerfile Best Practices

### Multi-Stage Builds
Use multiple `FROM` instructions to separate build and runtime dependencies.

### Layer Optimization
- Order instructions from least to most frequently changing
- Combine related `RUN` commands to minimize layers
- Clean up temporary files in the same `RUN` command

### .dockerignore
Create comprehensive `.dockerignore`

### COPY Strategy
- Copy dependency files first for better caching
- Copy specific paths, not entire directories when possible
- Use `.dockerignore` to exclude unnecessary files

### CMD & ENTRYPOINT
- Use exec form: `["command", "arg1", "arg2"]`
- Use `ENTRYPOINT` for executable, `CMD` for default arguments
- Prefer exec form for better signal handling

### Environment Variables
- Use `ENV` for defaults, allow runtime overrides
- Never hardcode secrets in Dockerfile
- Validate required variables at application startup

## Security Best Practices

### Non-Root User
- Always create and use dedicated non-root user to minimize privilege escalation risks.
- Set proper file permissions
- Document exposed ports with `EXPOSE`

### Minimal Images
- Use `alpine`
- Fewer packages = fewer vulnerabilities

### No Secrets in Layers
- Never include secrets, keys, or credentials in image layers
- Use runtime secrets management (Kubernetes Secrets, Docker Secrets, Vault)
- Scan images for accidentally included secrets

### Health Checks
Define `HEALTHCHECK` for liveness and readiness probes.

## Runtime Best Practices

### Resource Limits
Set CPU and memory limits to prevent resource exhaustion.

### Logging
- Log to `STDOUT`/`STDERR`

### Persistent Storage
- Use named volumes or Persistent Volumes for stateful data
- Never store data in container's writable layer
- Implement backup strategies

### Networking
- Create custom networks for isolation
- Define network policies
- Use service discovery mechanisms
