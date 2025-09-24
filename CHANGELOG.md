# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Tests for string replacements across supported formats (JSON, TypeScript, Properties).
- `codecov.yml` configuration file for proper code coverage reporting.
- Coverage file verification step in CI pipeline for debugging.

### Changed

- Renamed `BooleanSetting` to `InternalSetting` across the codebase.

### Fixed

- Codecov integration issues in GitHub Actions workflow.
- Updated codecov action from `file:` to `files:` parameter (v4 compatibility).
- Added verbose logging to codecov upload for better debugging.

### Technical Details

- End-to-end `erstatte.json` workflow test for string replacements.
- Enhanced CI pipeline with coverage file existence verification.
- Configured codecov with 70% coverage target and proper ignore patterns.

## [1.0.1] - 2025-09-10

### Added

- Initial release of erstatte configuration replacement tool
- Support for JSON, TypeScript, and Properties file formats
- Configuration-driven replacement via erstatte.json
- Cross-platform binary builds (Linux, Windows, macOS)
- Comprehensive test suite with 76.7% coverage
- GitHub Actions CI/CD pipeline
- Automated security scanning and code quality checks
- Performance benchmarks and monitoring
- Multi-version Go testing (1.21, 1.22, 1.23)

### Features

- **Multi-format support**: JSON, TypeScript (.ts), and Properties files
- **Type-aware replacements**: Intelligent handling of booleans, numbers, and strings
- **Package-based architecture**: Clean separation of concerns
- **Extensive testing**: Unit tests, integration tests, and benchmarks
- **Development tools**: Build scripts for Windows (PowerShell) and Unix (Make)

### Technical Details

- Written in Go with modular package structure
- Uses Cobra CLI framework for command-line interface
- Implements regex-based TypeScript parsing for type-aware replacements
- JSON parsing with structured object manipulation
- Properties file parsing with comment preservation

### Documentation

- Comprehensive README with usage examples
- Detailed test documentation (README_TESTS.md)
- GitHub Actions documentation (README_GITHUB_ACTIONS.md)
- Code coverage reports and benchmarks

### CI/CD

- Automated testing on multiple Go versions
- Cross-platform builds for all major operating systems
- Security vulnerability scanning with Gosec and Nancy
- Code quality checks with golangci-lint
- Automatic releases with GitHub Actions
- Coverage reporting with Codecov integration

## [1.0.0] - 2025-09-09

### Added

- Initial public release
- Core functionality for configuration file replacement
- Support for JSON, TypeScript, and Properties formats
- Command-line interface
- Comprehensive documentation and examples

[Unreleased]: https://github.com/johanneslosch/erstatte/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/johanneslosch/erstatte/releases/tag/v1.0.0
