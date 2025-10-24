# Script para iniciar todos los servicios con Docker Compose

Write-Host "Iniciando servicios con Docker Compose..." -ForegroundColor Green
Write-Host ""

# Verificar si Docker está ejecutándose
$dockerRunning = docker info 2>$null
if (-not $dockerRunning) {
    Write-Host "Error: Docker no esta ejecutandose." -ForegroundColor Red
    Write-Host "Por favor, inicia Docker Desktop e intenta de nuevo." -ForegroundColor Yellow
    exit 1
}

# Detener servicios previos si existen
Write-Host "Deteniendo servicios previos (si existen)..." -ForegroundColor Yellow
docker-compose down 2>$null

# Iniciar servicios
Write-Host "Iniciando servicios..." -ForegroundColor Cyan
docker-compose up -d

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "Servicios iniciados correctamente!" -ForegroundColor Green
    Write-Host ""
    Write-Host "CockroachDB Admin UI: http://localhost:8080" -ForegroundColor Cyan
    Write-Host "Database URL: postgresql://root@localhost:26257/stocks?sslmode=disable" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "Esperando 5 segundos para que los servicios esten listos..." -ForegroundColor Yellow
    Start-Sleep -Seconds 5
    
    Write-Host ""
    Write-Host "Ahora ejecuta el script de inicializacion:" -ForegroundColor Magenta
    Write-Host "   .\scripts\init-db.ps1" -ForegroundColor White
} else {
    Write-Host ""
    Write-Host "Error al iniciar los servicios." -ForegroundColor Red
    Write-Host "Verifica los logs con: docker-compose logs" -ForegroundColor Yellow
}
