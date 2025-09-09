# PowerShell Build Script for erstatte
param([string]$Command = "help")

function TestAll {
    Write-Host "🧪 Running all tests..." -ForegroundColor Cyan
    go test -v ./tests/unit/...
    go test -v ./tests/integration/...
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ All tests successful" -ForegroundColor Green
    }
}

function TestUnit {
    Write-Host "🧪 Running unit tests..." -ForegroundColor Cyan
    go test -v ./tests/unit/...
}

function BuildBinary {
    Write-Host "🏗️  Building binary..." -ForegroundColor Cyan
    if (!(Test-Path "dist")) { New-Item -ItemType Directory -Name "dist" }
    go build -o dist/erstatte.exe ./
}

function ShowHelp {
    Write-Host "Available commands:" -ForegroundColor Yellow
    Write-Host "  test      - All tests" -ForegroundColor White
    Write-Host "  unit      - Unit tests" -ForegroundColor White
    Write-Host "  build     - Build binary" -ForegroundColor White
    Write-Host "  help      - This help" -ForegroundColor White
}

switch ($Command.ToLower()) {
    "test" { TestAll }
    "unit" { TestUnit }
    "build" { BuildBinary }
    "help" { ShowHelp }
    default { ShowHelp }
}
