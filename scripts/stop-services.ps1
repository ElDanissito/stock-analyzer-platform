# Script para detener todos los servicios

Write-Host "🛑 Deteniendo servicios..." -ForegroundColor Yellow

docker-compose down

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "✅ Servicios detenidos correctamente!" -ForegroundColor Green
} else {
    Write-Host ""
    Write-Host "❌ Error al detener los servicios." -ForegroundColor Red
}
