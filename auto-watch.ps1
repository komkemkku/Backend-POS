# Auto-Watch and Deploy Script
# จะตรวจสอบการเปลี่ยนแปลงไฟล์อัตโนมัติและทำการ deploy

param(
    [int]$IntervalSeconds = 30  # ตรวจสอบทุก 30 วินาที
)

Write-Host "🔍 Starting auto-watch for Backend POS..." -ForegroundColor Green
Write-Host "⏱️  Checking for changes every $IntervalSeconds seconds" -ForegroundColor Yellow
Write-Host "🛑 Press Ctrl+C to stop watching" -ForegroundColor Red
Write-Host ""

$lastCommit = git rev-parse HEAD

while ($true) {
    Start-Sleep -Seconds $IntervalSeconds
    
    # ตรวจสอบว่ามีการเปลี่ยนแปลงหรือไม่
    $changes = git status --porcelain
    
    if ($changes.Length -gt 0) {
        Write-Host "🔄 Changes detected! Running auto-deploy..." -ForegroundColor Yellow
        
        # รันสคริปต์ auto-deploy
        & "$PSScriptRoot\auto-deploy.ps1"
        
        Write-Host "✅ Auto-deploy completed. Resuming watch..." -ForegroundColor Green
        Write-Host ""
        
        $lastCommit = git rev-parse HEAD
    } else {
        Write-Host "⏰ $(Get-Date -Format 'HH:mm:ss') - No changes detected" -ForegroundColor Gray
    }
}
