@echo off
echo 🚀 Auto-Deploy for Backend POS
echo.
powershell -ExecutionPolicy Bypass -File "%~dp0auto-deploy.ps1"
echo.
pause
