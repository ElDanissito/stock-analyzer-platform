# Script PowerShell para inicializar la base de datos CockroachDB

Write-Host "Inicializando base de datos CockroachDB..." -ForegroundColor Green

# Esperar a que CockroachDB este listo
Write-Host "Esperando a que CockroachDB este listo..." -ForegroundColor Yellow
Start-Sleep -Seconds 5

# Crear la base de datos
Write-Host "Creando base de datos 'stocks'..." -ForegroundColor Cyan
docker exec -it stock-analyzer-cockroachdb ./cockroach sql --insecure -e "CREATE DATABASE IF NOT EXISTS stocks;"

# Verificar que la base de datos fue creada
Write-Host "Verificando base de datos..." -ForegroundColor Cyan
docker exec -it stock-analyzer-cockroachdb ./cockroach sql --insecure -e "SHOW DATABASES;"

Write-Host ""
Write-Host "Base de datos inicializada correctamente!" -ForegroundColor Green
Write-Host ""
Write-Host "Puedes acceder al Admin UI en: http://localhost:8080" -ForegroundColor Yellow
Write-Host "Connection string: postgresql://root@localhost:26257/stocks?sslmode=disable" -ForegroundColor Yellow
Write-Host ""
Write-Host "Tip: Tambien puedes usar pgAdmin en http://localhost:5050" -ForegroundColor Magenta
Write-Host "   Email: admin@stock-analyzer.com" -ForegroundColor Magenta
Write-Host "   Password: admin" -ForegroundColor Magenta
