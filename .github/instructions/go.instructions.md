---
description: "Instructions for Go code generation"
applyTo: "**/*.go,**/go.mod,**/go.sum"
---

# Go Instructions

## Copyright Header

Each file must include this copyright header at the top using the current year:

```go
// Restricted - Copyright (C) Siemens Healthcare GmbH/Siemens Medical Solutions USA, Inc., [current year here]. All rights reserved.
```

## General Instructions

- Write simple, idiomatic Go code; favor clarity over cleverness
- Keep the happy path left-aligned; return early to reduce nesting
- Make the zero value useful
- Use Go modules and leverage the standard library
- Write comments in English; avoid emoji
- Prefer established codebase design patterns over narrowly minimizing the diff when both approaches are viable
- If following the existing design would be harmful, unclear, or materially more effort, stop and ask before diverging from it

## Naming Conventions

- **Packages**: Lowercase, single-word, singular (avoid `util`, `common`, `base`)
- **Variables/Functions**: mixedCaps or MixedCaps; exported names start with capital letter
- **Interfaces**: -er suffix when possible (e.g., `Reader`, `Writer`)
- **Constants**: MixedCaps for exported, mixedCaps for unexported
- **No abbreviations**: Use full words in identifiers (e.g., `subscriptionMutex` not `subscriptionMu`, `context` not `ctx` in struct fields)
- **Descriptive names for long-lived variables**: Non-temporary variables with a large scope (e.g., wired in `main`, stored in structs, passed across function boundaries) must have descriptive names that convey their purpose. Avoid generic names like `mgr`, `svc`, `cl`, or `srv`; prefer `controllerManager`, `emailService`, `restClient`, or `httpServer`.

### Package Declaration Rules (CRITICAL)

- **NEVER duplicate `package` declarations** - each file has exactly ONE
- When editing existing `.go` files: PRESERVE existing `package` line
- When creating new `.go` files: Use SAME package name as other files in that directory
- Before writing code, verify no existing `package` declaration exists

## Code Style

- Always use `gofmt` and `goimports`
- Strive for self-documenting code; comment only complex logic or business rules
- Prefer functions of 40 lines or fewer. If a longer function is necessary, add a short nearby justification for why splitting it further would reduce clarity.
- Check errors immediately; wrap with context using `fmt.Errorf` with `%w`
- Error messages: lowercase, no ending punctuation
- Exported symbols must be documented, starting with the symbol name

## Type Safety

- Use `any` instead of `interface{}` (Go 1.18+)
- Use pointer receivers for large structs or when modifying the receiver
- Accept interfaces, return concrete types
- Define interfaces close to usage, keep small (1-3 methods)

## Concurrency

- Document and provide cleanup for library goroutines; prefer letting callers control concurrency
- Always know how a goroutine exits; avoid leaks
- Close channels from sender side
- Use channels for communication, mutexes for protecting state
- WaitGroup:
  - `go >= 1.25`: Use `wg.Go(task)` method
  - `go < 1.25`: Use `Add`/`Done` pattern

## API Design

- **HTTP Router**: Use `net/http` `ServeMux` (Go 1.22+ has pattern-based routing)
- Validate input; handle errors gracefully with appropriate status codes
- Use struct tags for JSON; pointers for optional fields

### HTTP Clients

- Keep client struct for configuration only (base URL, `*http.Client`, auth); no per-request state
- Never cache `*http.Request` in client struct; construct fresh request per call
- Accept `context.Context` and parameters; build request locally
- Set headers on request instance; always close response bodies (`defer resp.Body.Close()`)

## I/O: Readers and Buffers

- `io.Reader` streams are consumable once; buffer if reuse needed (`io.ReadAll` → `bytes.NewReader`)
- HTTP request bodies: Don't reuse consumed `req.Body`; set `req.GetBody` for retries
- Use `io.TeeReader` to duplicate streams while reading
- `io.Pipe` for streaming: write in separate goroutine; always close writer; **writes must be sequential, not concurrent**
- Streaming multipart: `pr, pw := io.Pipe()`; `mw := multipart.NewWriter(pw)`; write all parts in order; close `mw` then `pw`

## Performance

- Minimize allocations in hot paths; preallocate slices when size known
- Use `sync.Pool` for reusable objects
- Profile before optimizing; focus on algorithmic improvements

## Testing

- Use table-driven tests with `t.Run` for subtests
- Name: `Test_functionName_scenario`
- Mark helpers with `t.Helper()`; clean up with `t.Cleanup()`
- **Every new feature, function, or behavior change must include corresponding unit tests**
- Place tests in the same package as the code under test (white-box testing)
- Cover the happy path, edge cases, and error paths

## Security

- Validate all external input; sanitize for SQL/HTML/shell contexts
- Use standard library crypto; `crypto/rand` for random numbers
- Hash passwords with bcrypt, scrypt, or argon2

## Tools

- Essential: `go fmt`, `go vet`, `golangci-lint`, `go test`, `go mod tidy`
- Run tests before committing; use pre-commit hooks

## Common Pitfalls

- Not checking errors or closing resources
- Race conditions and goroutine leaks
- Modifying maps concurrently
- Understanding nil interfaces vs nil pointers
- Over-using `any`; prefer specific types or constrained generics
- **Duplicate `package` declarations** - always verify before adding
