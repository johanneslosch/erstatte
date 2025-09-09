# Contributing to erstatte

Thank you for your interest in contributing to erstatte! This document provides guidelines and information for contributors.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Submitting Changes](#submitting-changes)
- [Code Style](#code-style)
- [Reporting Issues](#reporting-issues)

## 📜 Code of Conduct

This project adheres to a code of conduct. By participating, you are expected to uphold this code. Please be respectful and professional in all interactions.

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or later
- Git
- Basic knowledge of Go programming

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:

   ```bash
   git clone https://github.com/YOUR_USERNAME/erstatte.git
   cd erstatte
   ```

3. Add the original repository as upstream:
   ```bash
   git remote add upstream https://github.com/johanneslosch/erstatte.git
   ```

## 🛠️ Development Setup

### Install Dependencies

```bash
go mod download
go mod verify
```

### Verify Setup

```bash
# Run tests to ensure everything works
go test ./tests/...

# Build the binary
go build -o erstatte ./
```

### Using Build Scripts

#### Windows

```powershell
.\build-simple.ps1 test    # Run tests
.\build-simple.ps1 build   # Build binary
```

#### Unix

```bash
make test          # Run tests
make build         # Build binary
make dev          # Development setup
```

## 🔧 Making Changes

### Branch Naming

Create a descriptive branch name:

- `feature/add-yaml-support`
- `bugfix/fix-json-parsing`
- `docs/update-readme`
- `refactor/improve-error-handling`

### Commit Guidelines

Use clear, descriptive commit messages:

```bash
# Good
git commit -m "Add support for YAML configuration files"
git commit -m "Fix TypeScript boolean parsing edge case"

# Bad
git commit -m "Fix bug"
git commit -m "Update code"
```

### Keep Changes Focused

- One feature/fix per pull request
- Keep pull requests reasonably sized
- Include tests for new functionality

## 🧪 Testing

### Running Tests

```bash
# All tests
go test ./tests/...

# Specific test categories
go test ./tests/unit/...
go test ./tests/integration/...

# With coverage
go test -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out -o coverage.html
```

### Writing Tests

#### Test Organization

- **Unit tests**: `./tests/unit/`
- **Integration tests**: `./tests/integration/`
- **Benchmarks**: `./tests/benchmarks/`

#### Test Guidelines

- Test file names should end with `_test.go`
- Use table-driven tests when appropriate
- Include both positive and negative test cases
- Mock external dependencies
- Aim for high test coverage

#### Example Test Structure

```go
func TestParseAndReplace_JSON(t *testing.T) {
    // Setup
    tempDir := t.TempDir()
    testFile := filepath.Join(tempDir, "test.json")

    // Test data
    testData := map[string]interface{}{
        "key": "value",
    }

    // Execute
    err := parser.ParseAndReplace(testFile, replacements)

    // Assert
    if err != nil {
        t.Fatalf("Expected no error, got: %v", err)
    }

    // Verify results
    // ...
}
```

### Test Requirements

- All new features must include tests
- Bug fixes should include regression tests
- Maintain or improve test coverage
- All tests must pass before submitting

## 📝 Submitting Changes

### Before Submitting

1. **Update from upstream**:

   ```bash
   git fetch upstream
   git rebase upstream/master
   ```

2. **Run the full test suite**:

   ```bash
   go test ./tests/...
   ```

3. **Run linting**:

   ```bash
   golangci-lint run
   go vet ./...
   go fmt ./...
   ```

4. **Update documentation** if needed

### Pull Request Process

1. **Create a pull request** from your feature branch to the main repository's `master` branch

2. **Fill out the PR template** with:

   - Clear description of changes
   - Motivation for the changes
   - Testing performed
   - Breaking changes (if any)

3. **Ensure CI passes**:

   - All tests pass
   - Code quality checks pass
   - Security scans pass

4. **Address review feedback** promptly

5. **Keep PR up to date** with the main branch

### PR Template

```markdown
## Description

Brief description of changes

## Type of Change

- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing

- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Manual testing performed

## Checklist

- [ ] Code follows project style guidelines
- [ ] Self-review performed
- [ ] Documentation updated
- [ ] Tests added for new functionality
```

## 🎨 Code Style

### Go Style Guidelines

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Use `goimports` for import organization
- Follow Go naming conventions
- Write clear, self-documenting code

### Code Organization

- Keep functions small and focused
- Use meaningful variable and function names
- Add comments for complex logic
- Organize imports: standard library, third-party, local

### Example

```go
// ParseAndReplace replaces values in the specified file based on the
// provided replacements map. It supports JSON, TypeScript, and Properties files.
func ParseAndReplace(filePath string, replacements map[string]string) error {
    ext := strings.ToLower(filepath.Ext(filePath))

    content, err := os.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("failed to read file %s: %w", filePath, err)
    }

    // Process based on file type
    switch ext {
    case ".json":
        return processJSON(filePath, content, replacements)
    case ".ts":
        return processTypeScript(filePath, content, replacements)
    default:
        return fmt.Errorf("unsupported file type: %s", ext)
    }
}
```

## 🐛 Reporting Issues

### Before Reporting

1. Check existing issues to avoid duplicates
2. Try the latest version
3. Provide minimal reproduction case

### Issue Template

```markdown
## Bug Report

**Describe the bug**
Clear description of the issue

**To Reproduce**
Steps to reproduce:

1. Create file with content...
2. Run command...
3. See error

**Expected Behavior**
What should happen

**Environment**

- OS: [e.g., Windows 10, Ubuntu 20.04]
- Go version: [e.g., 1.21.0]
- erstatte version: [e.g., 1.0.0]

**Additional Context**
Any other relevant information
```

### Feature Requests

```markdown
## Feature Request

**Problem Description**
What problem does this solve?

**Proposed Solution**
Describe your proposed solution

**Alternatives Considered**
Other solutions you've considered

**Additional Context**
Any other relevant information
```

## 📞 Getting Help

- **Questions**: Open a GitHub issue with the `question` label
- **Discussions**: Use GitHub Discussions for general questions
- **Security Issues**: Email maintainers directly (see README)

## 🏆 Recognition

Contributors will be recognized in:

- GitHub contributors list
- Release notes for significant contributions
- README acknowledgments section

Thank you for contributing to erstatte! 🎉
