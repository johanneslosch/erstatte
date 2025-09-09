# PowerShell Build Script for erstatte
# Executes the same commands as the GitHub Actions

param(
    [Parameter(Position=0)]
    [string]$Command = "help"
)

function Test-Unit {
    Write-Host "🧪 Running unit tests..." -ForegroundColor Cyan
    go test -v ./tests/unit/...
}

function Test-Integration {
    Write-Host "🔗 Running integration tests..." -ForegroundColor Cyan
    go test -v ./tests/integration/...
}

function Test-All {
    Write-Host "🧪 Running all tests..." -ForegroundColor Cyan
    Test-Unit
    Test-Integration
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ Alle Tests erfolgreich" -ForegroundColor Green
    }
}

function Get-Coverage {
    Write-Host "📊 Creating coverage report..." -ForegroundColor Cyan
    go test -coverprofile=coverage.out ./tests/...
    go tool cover -html=coverage.out -o coverage.html
    Write-Host "Coverage Report: coverage.html" -ForegroundColor Green
}

function Invoke-Benchmark {
    Write-Host "⚡ Running benchmarks..." -ForegroundColor Cyan
    go test -bench=. -benchmem ./tests/benchmarks/...
}

function Invoke-Lint {
    Write-Host "🔍 Code Linting..." -ForegroundColor Cyan
    if (Get-Command golangci-lint -ErrorAction SilentlyContinue) {
        golangci-lint run
    } else {
        Write-Host "⚠️  golangci-lint nicht installiert. Verwende go vet..." -ForegroundColor Yellow
    }
    go vet ./...
    $fmt_output = go fmt ./...
    if ($fmt_output) {
        Write-Host "⚠️  Code nicht formatiert:" -ForegroundColor Yellow
        $fmt_output
    }
}

function Format-Code {
    Write-Host "🎨 Code formatieren..." -ForegroundColor Cyan
    go fmt ./...
    if (Get-Command goimports -ErrorAction SilentlyContinue) {
        goimports -w .
    }
}

function Build-Binary {
    Write-Host "🏗️  Building binary..." -ForegroundColor Cyan
    if (!(Test-Path "dist")) {
        New-Item -ItemType Directory -Name "dist"
    }
    go build -o dist/erstatte.exe ./
}

function Build-All {
    Write-Host "🏗️  Building cross-platform binaries..." -ForegroundColor Cyan
    if (!(Test-Path "dist")) {
        New-Item -ItemType Directory -Name "dist"
    }
    
    # Linux
    $env:GOOS = "linux"; $env:GOARCH = "amd64"
    go build -o dist/erstatte-linux-amd64 ./
    
    $env:GOOS = "linux"; $env:GOARCH = "arm64"
    go build -o dist/erstatte-linux-arm64 ./
    
    # Windows
    $env:GOOS = "windows"; $env:GOARCH = "amd64"
    go build -o dist/erstatte-windows-amd64.exe ./
    
    # macOS
    $env:GOOS = "darwin"; $env:GOARCH = "amd64"
    go build -o dist/erstatte-darwin-amd64 ./
    
    $env:GOOS = "darwin"; $env:GOARCH = "arm64"
    go build -o dist/erstatte-darwin-arm64 ./
    
    # Reset environment
    Remove-Item Env:GOOS -ErrorAction SilentlyContinue
    Remove-Item Env:GOARCH -ErrorAction SilentlyContinue
    
    Write-Host "✅ Binaries created in dist/" -ForegroundColor Green
}

function Install-Dependencies {
    Write-Host "📦 Dependencies installieren..." -ForegroundColor Cyan
    go mod download
    go mod verify
}

function Update-Dependencies {
    Write-Host "📦 Dependencies aktualisieren..." -ForegroundColor Cyan
    go get -u ./...
    go mod tidy
}

function Invoke-Security {
    Write-Host "🔒 Security Scan..." -ForegroundColor Cyan
    if (Get-Command gosec -ErrorAction SilentlyContinue) {
        gosec ./...
    } else {
        Write-Host "⚠️  gosec nicht installiert. Installation mit: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest" -ForegroundColor Yellow
    }
}

function Clear-Build {
    Write-Host "🧹 Cleaning up..." -ForegroundColor Cyan
    if (Test-Path "dist") {
        Remove-Item -Recurse -Force "dist"
    }
    if (Test-Path "coverage.out") {
        Remove-Item "coverage.out"
    }
    if (Test-Path "coverage.html") {
        Remove-Item "coverage.html"
    }
    go clean
}

function Setup-Development {
    Write-Host "🚀 Development Setup..." -ForegroundColor Cyan
    Install-Dependencies
    Format-Code
    Test-All
    Write-Host "✅ Development Setup abgeschlossen" -ForegroundColor Green
}

function Simulate-CI {
    Write-Host "🎯 CI Pipeline simulieren..." -ForegroundColor Cyan
    Install-Dependencies
    Test-All
    Invoke-Lint
    Invoke-Security
    Build-All
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ CI Pipeline erfolgreich" -ForegroundColor Green
    }
}

function Show-Help {
    Write-Host "Available commands:" -ForegroundColor Yellow
    Write-Host "  test          - Unit and integration tests" -ForegroundColor White
    Write-Host "  test-unit     - Unit tests only" -ForegroundColor White
    Write-Host "  test-integration - Integration tests only" -ForegroundColor White
    Write-Host "  coverage      - Create coverage report" -ForegroundColor White
    Write-Host "  benchmark     - Run benchmarks" -ForegroundColor White
    Write-Host "  lint          - Code linting" -ForegroundColor White
    Write-Host "  fmt           - Format code" -ForegroundColor White
    Write-Host "  build         - Binary for Windows" -ForegroundColor White
    Write-Host "  build-all     - Cross-platform binaries" -ForegroundColor White
    Write-Host "  deps          - Install dependencies" -ForegroundColor White
    Write-Host "  deps-update   - Update dependencies" -ForegroundColor White
    Write-Host "  security      - Security scan" -ForegroundColor White
    Write-Host "  clean         - Clean up" -ForegroundColor White
    Write-Host "  dev           - Development setup" -ForegroundColor White
    Write-Host "  ci            - Simulate CI pipeline" -ForegroundColor White
    Write-Host "  help          - Show this help" -ForegroundColor White
    Write-Host "" -ForegroundColor White
    Write-Host "Usage: .\build.ps1 <command>" -ForegroundColor Yellow
    Write-Host "Example: .\build.ps1 test" -ForegroundColor Yellow
}

# Hauptlogik
switch ($Command.ToLower()) {
    "test" { Test-All }
    "test-unit" { Test-Unit }
    "test-integration" { Test-Integration }
    "coverage" { Get-Coverage }
    "benchmark" { Invoke-Benchmark }
    "lint" { Invoke-Lint }
    "fmt" { Format-Code }
    "build" { Build-Binary }
    "build-all" { Build-All }
    "deps" { Install-Dependencies }
    "deps-update" { Update-Dependencies }
    "security" { Invoke-Security }
    "clean" { Clear-Build }
    "dev" { Setup-Development }
    "ci" { Simulate-CI }
    "help" { Show-Help }
    default { 
        Write-Host "❌ Unbekannter Befehl: $Command" -ForegroundColor Red
        Show-Help 
    }
}
