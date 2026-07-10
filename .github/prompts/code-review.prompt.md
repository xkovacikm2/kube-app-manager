---
description: "Prompt for doing a detailed code review"
agent: agent
model: "claude-sonnet-4.5"
tools: ["read/problems", "read/readFile", "search", "todo"]
name: "my code review"
---

## Code Review Prompt

**Precondition**: At least one code file must be in context. If no files are available, stop and ask the user to open or attach the files to refactor.
**Enforced Rule**: This code-review prompt is read-only — do not edit any files; only propose changes.

## Priority Levels

- **CRITICAL**: Security vulnerabilities, logic errors, breaking changes, data loss risks
- **IMPORTANT**: Code quality violations, missing tests, performance issues, architectural problems
- **SUGGESTION**: Readability, optimizations, minor best practices, documentation

## Review Principles

- Be specific with file/line references
- Explain WHY and impact, not just WHAT

## Quality Checks

### Clean Code

- Descriptive names (variables, functions, classes)
- Single Responsibility Principle
- DRY (no duplication)
- Functions < 30 lines, nesting < 4 levels
- Named constants instead of magic numbers

### Error Handling

- Proper error handling at appropriate levels
- Meaningful error messages
- No silent failures
- Validate inputs early
- Use appropriate error types

## Security Checks

- No secrets/credentials/Protected Health Information in code or logs
- Input validation and sanitization
- Avoid any form of injections
- Authentication/authorization checks
- Use established crypto libraries
- No vulnerable dependencies

## Testing Standards

- Critical paths and new functionality must have tests
- Descriptive test names (Given-When-Then)
- Arrange-Act-Assert structure
- Independent tests (no shared state)
- Specific assertions, test edge cases
- Mock external dependencies

## Performance Checks

- No N+1 queries (use joins/eager loading)
- Appropriate algorithm complexity
- Caching for expensive operations
- Resource cleanup (connections, files, streams)
- Pagination for large result sets
- No reflection in code

## Architecture Checks

- Separation of concerns
- Dependency direction (high→low level)
- Loose coupling, high cohesion
- Consistent with existing patterns

## Documentation Checks

- Public APIs documented (purpose, parameters, returns, exceptions)
- Complex logic has explanatory comments
- SDN updated for changes
