# erstatte - Configuration File Replacement Tool

**Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

erstatte is a Go CLI application that replaces values in JSON, TypeScript, and Properties files based on JSON configuration settings. It supports boolean, string, and number replacements with intelligent type handling for TypeScript files.

## Working Effectively

### Prerequisites and Setup
- **Go 1.21 or later required** (currently tested with Go 1.24.7)
- Repository uses Go modules with no external dependencies beyond standard library
- Cross-platform support: Linux, Windows, macOS (amd64 and arm64)

### Essential Commands and Timing

#### Dependencies (Instant)
```bash
go mod download    # <1 second - no external dependencies
go mod verify      # <1 second - verify module integrity
```

#### Building (Fast)
```bash
# Single platform build
go build -o erstatte ./         # 0.2 seconds
make build                      # 0.2 seconds - builds to dist/erstatte

# Cross-platform builds - NEVER CANCEL: Takes 32 seconds, set timeout to 60+ seconds
make build-all                  # Builds for Linux, Windows, macOS (both amd64/arm64)
```

#### Testing - NEVER CANCEL: Set timeout to 30+ seconds
```bash
# All tests (unit + integration + benchmarks)
go test ./tests/...             # 19 seconds total
make test                       # 0.2 seconds (cached results)

# Specific test categories
go test ./tests/unit/...        # Unit tests only
go test ./tests/integration/... # Integration tests only
make benchmark                  # 6 seconds - performance benchmarks with memory stats
```

#### Code Quality and Linting
```bash
# These work reliably
go vet ./...                    # 1.3 seconds - static analysis
go fmt ./...                    # 0.65 seconds - code formatting

# golangci-lint has version compatibility issues with Go 1.24.7
# Use go vet and go fmt instead for reliable linting
```

#### Coverage Reports
```bash
make coverage                   # 2.2 seconds - generates coverage.html report
go test -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out -o coverage.html
```

## Manual Validation Requirements

### ALWAYS test application functionality after making changes:

#### 1. TypeScript File Replacement Test
```bash
# Create test TypeScript file
cat > testfile.ts << 'EOF'
export const config = {
  debugMode: false,
  environment: "development",
  resetDatabase: false
};
EOF

# Create settings.json
cat > settings.json << 'EOF'
{
  "boolean": [
    {
      "filePath": "./testfile.ts",
      "field": "debugMode",
      "value": "true"
    },
    {
      "filePath": "./testfile.ts", 
      "field": "environment",
      "value": "production"
    }
  ]
}
EOF

# Run the tool and verify output
./erstatte
# Expected output: "./testfile.ts: debugMode = true replaced"
#                  "./testfile.ts: environment = production replaced"

# Verify the file contents changed correctly:
# debugMode should become: debugMode: true (boolean without quotes)
# environment should become: environment: "production" (string with quotes)
```

#### 2. JSON File Replacement Test
```bash
# Create test JSON file
cat > config.json << 'EOF'
{
  "debugMode": false,
  "port": 3000,
  "environment": "development"
}
EOF

# Update settings.json for JSON test
cat > settings.json << 'EOF'
{
  "boolean": [
    {
      "filePath": "./config.json",
      "field": "debugMode", 
      "value": "true"
    },
    {
      "filePath": "./config.json",
      "field": "environment",
      "value": "production"
    }
  ]
}
EOF

# Run and verify JSON replacement
./erstatte
# Verify: JSON values are replaced as strings (with quotes)
```

### Build Validation - CRITICAL
```bash
# Always build and test before committing changes
make test && make build
# Test the built binary with sample files
./dist/erstatte
```

## Project Structure and Navigation

### Core Directories
```
/pkg/settings/     # Settings management (JSON config loading/saving)
/pkg/parser/       # File parsing and replacement logic  
/tests/unit/       # Unit tests for individual components
/tests/integration/ # End-to-end workflow tests
/tests/benchmarks/ # Performance benchmarks
/examples/         # Example configuration files and test data
```

### Key Files
- `main.go` - Application entry point (loads settings.json, processes replacements)
- `Makefile` - Unix build commands (test, build, coverage, cross-compilation)
- `.golangci.yml` - Linting configuration (has compatibility issues with Go 1.24.7)
- `settings.json` - Default configuration file (created by user)

### Supported File Types
- **JSON files (`.json`)** - Replaces values as strings
- **TypeScript files (`.ts`)** - Intelligent type-aware replacement (booleans without quotes, strings with quotes)
- **Properties files (`.properties`)** - Key-value pair replacement

## Common Development Tasks

### Making Code Changes
1. **Always run tests first**: `make test` (validates current state)
2. **Make minimal changes** to achieve the goal
3. **Build immediately**: `go build -o erstatte ./`
4. **Manual validation**: Test with both TypeScript and JSON files as shown above
5. **Final validation**: `make test && go vet ./...`

### Adding New Features
1. Check existing tests in `/tests/unit/` and `/tests/integration/`
2. Add tests first, then implement feature
3. Use the benchmark tests in `/tests/benchmarks/` as performance guides
4. Follow existing code patterns in `/pkg/settings/` and `/pkg/parser/`

### Debugging Issues
- Test files are in `/examples/` directory (testfile.ts, example-settings.json)  
- Integration tests in `/tests/integration/` show complete workflows
- Use `go run ./` for quick testing during development
- Settings file must be named `settings.json` in current directory

## CI/CD and GitHub Actions

### Automated Workflows
- **Tests**: Runs on Go 1.21, 1.22, 1.23 (multi-version compatibility)
- **Code Quality**: golangci-lint, go vet, security scanning (gosec)
- **Builds**: Cross-platform binary creation for releases

### Pre-commit Validation  
```bash
# Simulate CI pipeline locally
go mod verify
go test ./tests/...      # NEVER CANCEL: 19 seconds
go vet ./...            # 1.3 seconds
go fmt ./...            # 0.65 seconds  
make build-all          # NEVER CANCEL: 32 seconds, set 60+ second timeout
```

## Performance Expectations (AMD EPYC 7763)
- JSON parsing: ~433μs per operation
- TypeScript parsing: ~425μs per operation  
- Settings loading: ~15μs per operation
- Settings saving: ~44μs per operation

## Troubleshooting

### Common Issues
- **golangci-lint fails**: Version compatibility issue with Go 1.24.7 - use `go vet` and `go fmt` instead
- **No settings.json**: Application requires `settings.json` in current directory
- **File not found**: Ensure file paths in settings.json are relative to execution directory
- **Cross-compilation slow**: Normal behavior - 32 seconds for 5 platform builds

### Validation Failures
If manual testing fails:
1. Check settings.json format matches examples
2. Verify file paths are correct relative to execution directory  
3. Test with known-good examples from `/examples/` directory
4. Check file permissions and write access

---

**CRITICAL REMINDERS:**
- **NEVER CANCEL builds or tests** - builds take 32 seconds, tests take 19 seconds
- **ALWAYS test manually** with both TypeScript and JSON files after changes
- **Use go vet/go fmt** instead of golangci-lint for Go 1.24.7 compatibility
- **Set 60+ second timeouts** for cross-platform builds
- **Set 30+ second timeouts** for test suites