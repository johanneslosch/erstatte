# erstatte

[![Tests](https://github.com/johanneslosch/erstatte/workflows/Tests/badge.svg)](https://github.com/johanneslosch/erstatte/actions/workflows/test.yml)
[![Code Quality](https://github.com/johanneslosch/erstatte/workflows/Code%20Quality/badge.svg)](https://github.com/johanneslosch/erstatte/actions/workflows/quality.yml)
[![Release](https://github.com/johanneslosch/erstatte/workflows/Release/badge.svg)](https://github.com/johanneslosch/erstatte/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/johanneslosch/erstatte)](https://goreportcard.com/report/github.com/johanneslosch/erstatte)
[![codecov](https://codecov.io/gh/johanneslosch/erstatte/branch/master/graph/badge.svg)](https://codecov.io/gh/johanneslosch/erstatte)

A powerful configuration file replacement tool written in Go that allows you to dynamically replace values in JSON, TypeScript, and Properties files based on configurable settings.

## 🚀 Features

- **Multi-format support**: JSON, TypeScript (.ts), and Properties files
- **Configuration-driven**: Define replacements via JSON settings file
- **Type-aware replacements**: Intelligent handling of booleans, numbers, and strings in TypeScript
- **Cross-platform**: Runs on Linux, Windows, and macOS
- **Fast and reliable**: Built with Go for performance and reliability

## 📦 Installation

### Download Binary

Download the latest binary from the [releases page](https://github.com/johanneslosch/erstatte/releases):

#### Linux/macOS

```bash
# Download and install (replace with latest version)
wget https://github.com/johanneslosch/erstatte/releases/download/v1.0.0/erstatte-linux-amd64.tar.gz
tar -xzf erstatte-linux-amd64.tar.gz
chmod +x erstatte-linux-amd64
sudo mv erstatte-linux-amd64 /usr/local/bin/erstatte
```

#### Windows

Download `erstatte-windows-amd64.zip` from the releases page and extract the executable.

### Build from Source

```bash
git clone https://github.com/johanneslosch/erstatte.git
cd erstatte
go build -o erstatte ./
```

## 🔧 Usage

### Basic Usage

1. Create a `settings.json` file (see [Configuration](#configuration) or check `examples/` directory)
2. Run the tool:

```bash
./erstatte

# Or specify a custom settings file
./erstatte --settings examples/example-settings.json
```

The tool will:

- Load settings from `settings.json`
- Apply all configured replacements to the specified files
- Display the results

### Configuration

Create a `settings.json` file in the same directory as the executable:

```json
{
  "boolean": [
    {
      "filePath": "./config.ts",
      "field": "debugMode",
      "value": "true"
    },
    {
      "filePath": "./database.json",
      "field": "enableLogging",
      "value": "false"
    },
    {
      "filePath": "./app.properties",
      "field": "production.mode",
      "value": "true"
    }
  ]
}
```

#### Configuration Fields

- **`filePath`**: Path to the target file (relative or absolute)
- **`field`**: The field/key to replace
- **`value`**: The new value to set

## 📁 Supported File Formats

### JSON Files (`.json`)

Replaces values in JSON objects:

```json
// Before
{
  "debugMode": false,
  "maxRetries": 3
}

// After (debugMode: "true", maxRetries: "10")
{
  "debugMode": "true",
  "maxRetries": "10"
}
```

### TypeScript Files (`.ts`)

Intelligently replaces values in TypeScript configuration objects:

```typescript
// Before
export const config = {
  debugMode: false,
  maxRetries: 3,
  serverUrl: "localhost",
};

// After (debugMode: "true", serverUrl: "production.com")
export const config = {
  debugMode: true, // Boolean without quotes
  maxRetries: 3, // Unchanged
  serverUrl: "production.com", // String with quotes
};
```

### Properties Files (`.properties`)

Replaces key-value pairs in properties files:

```properties
# Before
debug.enabled=false
server.port=8080

# After (debug.enabled: "true")
debug.enabled=true
server.port=8080
```

## 🏗️ Project Structure

```
erstatte/
├── pkg/                          # Core packages
│   ├── settings/                 # Settings management
│   │   └── settings.go
│   └── parser/                   # File parsing and replacement
│       └── parser.go
├── tests/                        # Test organization
│   ├── unit/                     # Unit tests
│   │   ├── settings/
│   │   └── parser/
│   ├── integration/              # Integration tests
│   └── benchmarks/               # Performance tests
├── examples/                     # Example files and configurations
│   ├── example-settings.json     # Example configuration
│   ├── testfile.ts              # Example TypeScript file
│   └── README.md                # Examples documentation
├── .github/workflows/            # GitHub Actions CI/CD
├── main.go                       # Main application
└── build-simple.ps1             # Build script for Windows
```

## 🧪 Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile)

### Running Tests

```bash
# All tests
go test ./tests/...

# Unit tests only
go test ./tests/unit/...

# Integration tests only
go test ./tests/integration/...

# With coverage
go test -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out -o coverage.html
```

### Using Build Scripts

#### Windows (PowerShell)

```powershell
.\build-simple.ps1 test    # Run tests
.\build-simple.ps1 build   # Build binary
.\build-simple.ps1 help    # Show all commands
```

#### Unix (Make)

```bash
make test          # Run tests
make build         # Build binary
make build-all     # Cross-platform builds
make help          # Show all targets
```

### Benchmarks

```bash
go test -bench=. ./tests/benchmarks/...
```

Performance results on AMD Ryzen 5 3500X:

- JSON parsing: ~1.9ms per operation
- TypeScript parsing: ~1.7ms per operation
- Settings loading: ~89μs per operation

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests for your changes
5. Run the test suite (`go test ./tests/...`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Code Quality

This project uses several tools to maintain code quality:

- **golangci-lint**: Static analysis and linting
- **gosec**: Security vulnerability scanning
- **gofmt**: Code formatting
- **go vet**: Go code analysis

All checks are automatically run in CI/CD pipeline.

## 🏷️ Versioning

This project uses [Semantic Versioning](https://semver.org/). Releases are automatically created when version tags are pushed.

To create a new release:

```bash
git tag v1.0.0
git push origin v1.0.0
```

## 📊 CI/CD

The project includes comprehensive GitHub Actions workflows:

- **Tests**: Multi-version Go testing (1.21, 1.22, 1.23)
- **Code Quality**: Linting, security scanning, dependency checks
- **Releases**: Automatic binary builds and GitHub releases
- **Cross-platform**: Linux, Windows, macOS support

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [GitHub Repository](https://github.com/johanneslosch/erstatte)
- [Issues](https://github.com/johanneslosch/erstatte/issues)
- [Releases](https://github.com/johanneslosch/erstatte/releases)
- [Documentation](README_TESTS.md) (Test organization)
- [GitHub Actions](README_GITHUB_ACTIONS.md) (CI/CD setup)

## 🙏 Acknowledgments

- Built with [Cobra CLI](https://github.com/spf13/cobra) framework
- Inspired by the need for dynamic configuration management
- Thanks to the Go community for excellent tooling

## 📈 Project Stats

- **Language**: Go
- **Test Coverage**: 76.7%
- **Tests**: 9 unit/integration tests + 4 benchmarks
- **Supported Platforms**: Linux, Windows, macOS
- **Supported Architectures**: amd64, arm64

---

<div align="center">
  <strong>Made with ❤️ and Go</strong>
</div>
