---
description: "Prompt for simple, single-file code refactoring with verification"
agent: agent
model: "claude-sonnet-4.5"
tools: ["read/problems", "read/readFile", "edit/editFiles"]
name: "my refactor"
---

## Refactor Prompt

**Precondition**: Exactly one code file must be in context. If no file or multiple files are open, stop and ask the user to focus on a single file.

**Scope**: This prompt handles simple, single-file refactorings only. For complex multi-file refactoring, manual approach is recommended.

**Refactor Goal**: Use the user's natural language input to determine what to refactor. If the user provided no specific guidance, default to improving code structure, readability, and maintainability within the file.

## Execution Protocol

Follow these steps strictly. **Stop and report** if any step fails verification.

### Step 1: Analysis (read-only)

1. Review the current file to understand its structure and purpose
2. Identify the specific code section to refactor (use selection if provided)
3. Check for existing tests in the same file (if applicable)
4. **Report**: Describe what will be changed and why before proceeding

### Step 2: Apply Refactoring

1. Make the refactoring changes within the single file
2. Preserve existing public API and behavior
3. Keep changes minimal and focused
4. **Verify**: Run `read/problems` — zero new errors before continuing

### Step 3: Documentation

1. Update comments only if signatures or behavior changed
2. Ensure code is self-documenting
3. Add brief explanatory comments only for non-obvious logic

### Step 4: Final Verification

1. Run `read/problems` — confirm zero errors
2. Verify no unintended changes were made

## Constraints

- **Preserve behavior**: Refactored code must produce identical results
- **No partial edits**: Complete the refactoring fully or revert
- **Minimal scope**: Change only what's necessary for the refactoring goal

## On Failure

If refactoring cannot be completed:

1. Report what failed and why
2. Revert any partial changes
3. Suggest what the user should do manually
