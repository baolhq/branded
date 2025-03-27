# Set paths
$SRC = "main.go"
$OUT_DIR = "dist"

# Detect system architecture
$arch = if ([System.Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
$OUT = "$OUT_DIR/branded-$arch.exe"

# Ensure the output directory exists
if (!(Test-Path $OUT_DIR)) {
    New-Item -ItemType Directory -Path $OUT_DIR | Out-Null
}

# Build for the detected architecture
Write-Host "Detected system architecture: Windows $arch-bit"
Write-Host "Building for Windows $arch-bit..."

$env:GOOS = "windows"
$env:GOARCH = $arch
go build -o $OUT $SRC

# Check if build was successful
if ($?) {
    Write-Host "Build successful: $OUT" -ForegroundColor Green
} else {
    Write-Host "Build failed." -ForegroundColor Red
    exit 1
}
