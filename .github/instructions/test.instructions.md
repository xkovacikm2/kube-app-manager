---
description: "Instructions for test generation."
applyTo: "**/Test/**"
---
# Test Instructions
- Copy existing style from nearby files for test method names and capitalization

# Test Types
- Based on the type of the test apply the following rules

## Unit Tests
- Use Arrange, Act, Assert pattern, or if necessary combine Act and Assert into Act&Assert block
- Each test should focus on a single behavior or scenario
- The test should be self-contained, independent of other tests and leave the environment unchanged
- Use meaningful test names that convey the purpose of the test

## Integration Tests
- Each test should focus on a single behavior or scenario
- The test cannot take a long time to execute
- The test should be self-contained, independent of other tests and leave the environment unchanged
- Use meaningful test names starting with TC, that convey the purpose of the test
