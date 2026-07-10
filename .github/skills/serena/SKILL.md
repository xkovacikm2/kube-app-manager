---
name: serena
description: >-
  Semantic code navigation and editing using the Serena MCP server. Use when:
  finding Go symbols, tracing reconciler and client call paths, locating
  references, renaming symbols across internal packages, inserting or
  replacing code at the symbol level, or performing IDE-like navigation in
  this repository. Prefer these tools over full-file reads and text searches.
  Invoke for tasks like: find all usages of a Reconciler method, inspect
  SSE or REST client flows, refactor a package-level function, or
  understand the public surface of a Go type.
---

## Serena In This Repo

Serena provides IDE-grade symbol navigation and editing through the MCP
protocol. In this workspace the server starts automatically with the `ide`
context scoped to `${workspaceFolder}`, so no manual project activation is
required.

## When to Use Serena Tools

| Goal | Serena Tool | Instead of |
|------|-------------|------------|
| Find a function, method, struct, interface, or constant | `find_symbol` | `grep_search` plus guessing file paths |
| Get a quick map of a Go file | `get_symbols_overview` | Reading the entire file |
| Read a single symbol body | `find_symbol` with `include_body=true` | `read_file` on large files |
| See where a symbol is used | `find_referencing_symbols` | Manual cross-file analysis |
| Replace an entire function, method, or type definition | `replace_symbol_body` | `replace_content` with fragile anchors |
| Insert a new function, method, or type near existing code | `insert_after_symbol` or `insert_before_symbol` | Manual patching with large context |
| Rename a symbol across the repo | `rename_symbol` | Text search and replace |
| Make a small non-symbolic edit | `replace_content` | Rewriting a full symbol |
| Check onboarding and repo notes | `check_onboarding_performed`, `list_memories`, `read_memory` | N/A |

## Workflow for Code Changes

1. Start with `get_symbols_overview` when a Go file is unfamiliar,
   especially in `internal/controller` or the client packages.
2. Use `find_symbol` with `relative_path` to narrow the search and avoid
   collisions between repeated names such as `Client` or `Sender`.
3. Read only the symbol you need with `find_symbol` and
   `include_body=true`. For example, inspect `EmailTemplateReconciler/Reconcile`
   or `Client/GetResource` without loading the full file.
4. Use `find_referencing_symbols` before changing exported functions or
   shared helpers so dependent packages stay consistent.
5. Prefer `replace_symbol_body` for full symbol edits and `replace_content`
   for small line-level changes inside a larger function.
6. Validate with `go test` on the relevant package or `go test ./...` for
   wider changes. Run `make lint` when the edit affects multiple packages or
   public behavior.

## Codebase Conventions

This workspace is a Go email microservice that watches Kubernetes
EmailTemplate CRDs and sends emails via SMTP when matching SSE events
arrive from scida-rest-storage.

Key landmarks:

- `cmd/scida-email/main.go` wires the application: config, SMTP sender,
  watch client, SSE client, controller-runtime manager, and reconciler.
- `internal/controller` contains the `EmailTemplateReconciler` which manages
  SSE subscriptions per EmailTemplate and orchestrates the fetch → extract →
  render → send pipeline.
- `internal/email` contains the `Sender` interface and `SMTPSender`
  implementation using crypto/tls on port 465.
- `internal/sse` contains the SSE client that subscribes to
  scida-rest-storage event streams with exponential backoff reconnection.
- `internal/template` contains Go template rendering for email
  subject/body/recipients and JSONPath value extraction via ojg.
- `internal/watchclient` contains the polling client for fetching watched
  resources from scida-rest-storage.
- `internal/config` loads SMTP and service configuration from environment
  variables.
- `pkg/api/v1alpha1` defines the `EmailTemplate` CRD types under
  `email.scida.io/v1alpha1`.
- Tests live next to implementation files with `_test.go` suffixes.

## Tips for Efficiency

- Pass `relative_path` whenever possible. This repo has client-like types in
  multiple packages (`sse`, `watchclient`), so package scoping matters.
- Use name paths like `EmailTemplateReconciler/Reconcile`,
  `Client/Subscribe`, `Client/GetResource`, or `SMTPSender/Send` to jump
  directly to the right symbol.
- Combine `get_symbols_overview` with targeted `find_symbol` calls instead of
  reading whole Go files.
- Check Serena memories after onboarding for repo-specific commands,
  conventions, and completion steps.
