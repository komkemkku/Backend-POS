# สคริปต์สำหรับลบไฟล์ที่ไม่จำเป็น
# Clean unnecessary files script

Write-Host "Cleaning unnecessary files..." -ForegroundColor Green

# ลบไฟล์ build artifacts
$filesToDelete = @(
    "*.exe",
    "*.dll", 
    "*.so",
    "*.dylib",
    "main",
    "pos-server",
    "Backend-POS",
    "Backend-POS.exe"
)

foreach ($pattern in $filesToDelete) {
    $files = Get-ChildItem -Path . -Name $pattern -ErrorAction SilentlyContinue
    if ($files) {
        foreach ($file in $files) {
            Write-Host "Removing: $file" -ForegroundColor Yellow
            Remove-Item $file -Force -ErrorAction SilentlyContinue
        }
    }
}

# ลบ vendor directory หากมี
if (Test-Path "vendor") {
    Write-Host "Removing vendor directory..." -ForegroundColor Yellow
    Remove-Item -Recurse -Force "vendor" -ErrorAction SilentlyContinue
}

# ลบ logs หากมี
if (Test-Path "logs") {
    Write-Host "Removing logs directory..." -ForegroundColor Yellow
    Remove-Item -Recurse -Force "logs" -ErrorAction SilentlyContinue
}

# ลบไฟล์ log
$logFiles = Get-ChildItem -Path . -Name "*.log" -ErrorAction SilentlyContinue
if ($logFiles) {
    foreach ($logFile in $logFiles) {
        Write-Host "Removing log file: $logFile" -ForegroundColor Yellow
        Remove-Item $logFile -Force -ErrorAction SilentlyContinue
    }
}

# ลบ temporary files
if (Test-Path "tmp") {
    Write-Host "Removing tmp directory..." -ForegroundColor Yellow
    Remove-Item -Recurse -Force "tmp" -ErrorAction SilentlyContinue
}

if (Test-Path "temp") {
    Write-Host "Removing temp directory..." -ForegroundColor Yellow
    Remove-Item -Recurse -Force "temp" -ErrorAction SilentlyContinue
}

# Clean Go cache
Write-Host "Cleaning Go build cache..." -ForegroundColor Blue
go clean -cache -modcache -i -r 2>$null

Write-Host "Cleanup completed!" -ForegroundColor Green
Write-Host "Current directory size:"
$size = (Get-ChildItem -Recurse | Measure-Object -Property Length -Sum).Sum / 1MB
Write-Host "$([math]::Round($size, 2)) MB" -ForegroundColor Cyan
