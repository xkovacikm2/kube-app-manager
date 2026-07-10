---
description: "Prompt for reviewing merge via git diff with code quality and merge-specific checks"
agent: agent
model: "claude-sonnet-4.5"
tools:
  [
    "execute",
    "read/problems",
    "read/readFile",
    "read/terminalSelection",
    "read/terminalLastCommand",
    "search",
    "com.microsoft/azure/search",
    "todo",
  ]
name: "my merge review"
---

## Merge Review Prompt

**Precondition**: Git repository in context. Review is read-only — propose changes only, do not edit files.

Review a possible merge by analyzing `git diff` output. Apply all code-review checks plus merge-specific validation.

## Diff Retrieval

1. **Get current branch**: Execute `git branch --show-current` to identify the working branch
2. **Default comparison**: `git diff Main..<current-branch>` (current branch vs Main)
3. **User overrides**:
   - Custom base: `/cpdm-merge-review base:Feature/my-feature` → `git diff Feature/my-feature..<current-branch>`
   - Custom current: `/cpdm-merge-review current:Dev` → `git diff Main..Dev`
   - Both: `/cpdm-merge-review base:Feature/hotfix current:Feature/bugfix` → `git diff Feature/hotfix..Feature/bugfix`
4. **Validate merge size**: If diff exceeds ~500 lines changed, warn and ask for scope reduction. Proceed with available diff.

## Parse Diff Output

Extract from unified diff format:

- Changed files (A/M/D = added/modified/deleted)
- Line ranges (before/after line numbers)
- Context (3 lines before/after each chunk)

## Review Checks

- All Code-Review Checks Apply

### Merge-Specific Checks

- **Breaking changes**: Identified and documented?
- **Incomplete work**: TODOs, FIXMEs, console logs left behind?
- **Test coverage**: New code paths have tests?
- **Backward compatibility**: Will existing users/APIs break?
- **Documentation updates**: README, API docs, CHANGELOG, migration guides?
- **Dependencies**: New deps added? Justified? Version constraints reasonable?
- **File organization**: Respects existing structure and patterns?

## Priority Levels

- **CRITICAL**: Security, data loss, breaking API changes, incomplete implementation
- **IMPORTANT**: Code quality, missing tests, performance, architectural violations
- **SUGGESTION**: Readability, minor optimizations, documentation improvements

## Output Format

Group findings by:

1. **File** (from diff)
2. **Line range** (with context snippet from diff)
3. **Priority level**
4. **Finding type** (Clean Code, Security, Testing, etc.)
5. **Description** — explain WHY and impact
6. **Suggestion** — specific fix with code example

**Summary**:

- Ready to merge / Request changes / Needs discussion
- Total issues by priority level
- Estimated effort to address

## Instructions

1. Execute `git branch --show-current` to get the current branch
2. Parse user input for `base:` and `current:` overrides (defaults: base=Main, current=detected branch)
3. Execute `git diff <base>..<current>`
4. Parse diff to identify changed files and line ranges
5. For each hunk, read full file context via `read/readFile`
6. Apply all code-review checks + PR-specific checks
7. Output **only the top 15 findings** in concise format
8. Provide 3-5 line summary with merge readiness

The agent automatically retrieves, parses, and reviews all affected files.
