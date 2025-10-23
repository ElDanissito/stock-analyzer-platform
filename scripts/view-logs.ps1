# Script para ver los logs de los servicios

param(
    [string]$Service = "cockroachdb"
)

Write-Host "📋 Mostrando logs de $Service..." -ForegroundColor Cyan
Write-Host "Presiona Ctrl+C para salir" -ForegroundColor Yellow
Write-Host ""

docker-compose logs -f $Service
