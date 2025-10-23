# 🚀 Guía de Configuración - Stock Analyzer Platform

Esta guía te ayudará a configurar el proyecto desde cero.

## 📋 Prerequisitos

### Software Necesario

1. **Docker Desktop**
   - Descarga: https://www.docker.com/products/docker-desktop
   - Asegúrate de que esté ejecutándose

2. **Go 1.21+**
   - Descarga: https://go.dev/dl/
   - Verifica con: `go version`

3. **Node.js 18+**
   - Descarga: https://nodejs.org/
   - Verifica con: `node --version`

4. **Git**
   - Descarga: https://git-scm.com/
   - Verifica con: `git --version`

---

## 🐳 Paso 1: Configurar la Base de Datos

### Opción A: Usando Scripts PowerShell (Recomendado para Windows)

```powershell
# 1. Iniciar servicios (CockroachDB + pgAdmin)
.\scripts\start-services.ps1

# 2. Inicializar la base de datos
.\scripts\init-db.ps1

# Ver logs si hay problemas
.\scripts\view-logs.ps1
```

### Opción B: Comandos Manuales

```powershell
# Iniciar servicios
docker-compose up -d

# Esperar 5 segundos
Start-Sleep -Seconds 5

# Crear base de datos
docker exec -it stock-analyzer-cockroachdb ./cockroach sql --insecure -e "CREATE DATABASE IF NOT EXISTS stocks;"
```

### Verificar que funciona

- 🌐 **Admin UI**: http://localhost:8080
- 🔧 **pgAdmin**: http://localhost:5050
  - Email: `admin@stock-analyzer.com`
  - Password: `admin`

---

## ⚙️ Paso 2: Configurar el Backend

```powershell
# 1. Ir al directorio del backend
cd backend

# 2. Copiar el archivo de configuración
Copy-Item .env.example .env

# 3. Editar el archivo .env con tus credenciales
# (Abre backend\.env y actualiza los valores)

# 4. Descargar dependencias
go mod download
go mod tidy

# 5. Ejecutar migraciones
go run cmd/migrate/main.go

# 6. Iniciar el servidor
go run cmd/server/main.go
```

### Archivo `.env` del Backend

Edita `backend\.env`:

```env
PORT=8080
ENV=development
DATABASE_URL=postgresql://root@localhost:26257/stocks?sslmode=disable
STOCK_API_URL=https://api.karenai.click/swechallenge/list
STOCK_API_KEY=tu_api_key_aqui
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000
```

### Verificar que funciona

```powershell
# Test del endpoint de health
curl http://localhost:8080/health
```

---

## 🎨 Paso 3: Configurar el Frontend (Próximamente)

_Esta sección se completará cuando creemos el frontend._

---

## 🧪 Paso 4: Probar la Aplicación

### Sincronizar datos de la API

```powershell
# Método 1: Usando curl
curl -X POST http://localhost:8080/api/sync

# Método 2: Usando PowerShell
Invoke-RestMethod -Method POST -Uri "http://localhost:8080/api/sync"
```

### Ver los stocks

```powershell
# Listar stocks
curl http://localhost:8080/api/stocks

# Ver recomendaciones
curl http://localhost:8080/api/recommendations
```

---

## 📊 Comandos Útiles

### Docker

```powershell
# Ver servicios corriendo
docker-compose ps

# Ver logs
.\scripts\view-logs.ps1

# Detener servicios
.\scripts\stop-services.ps1

# Reiniciar servicios
docker-compose restart

# Eliminar todo (incluyendo datos)
docker-compose down -v
```

### Backend

```powershell
# Correr tests
cd backend
go test ./... -v

# Ver cobertura
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Formatear código
go fmt ./...
```

### Base de Datos

```powershell
# Conectar a CockroachDB CLI
docker exec -it stock-analyzer-cockroachdb ./cockroach sql --insecure --database=stocks

# Ver tablas
docker exec -it stock-analyzer-cockroachdb ./cockroach sql --insecure -e "USE stocks; SHOW TABLES;"

# Ver datos
docker exec -it stock-analyzer-cockroachdb ./cockroach sql --insecure -e "USE stocks; SELECT * FROM stocks LIMIT 10;"
```

---

## 🐛 Solución de Problemas

### Error: "Docker no está ejecutándose"

- Abre Docker Desktop
- Espera a que esté completamente iniciado (ícono verde)
- Intenta de nuevo

### Error: "Puerto 8080 ya en uso"

```powershell
# Opción 1: Detener el proceso que usa el puerto
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# Opción 2: Cambiar el puerto en docker-compose.yml y .env
```

### Error: "No se puede conectar a la base de datos"

```powershell
# Verificar que CockroachDB está corriendo
docker ps

# Ver logs
.\scripts\view-logs.ps1

# Reiniciar servicios
docker-compose restart
```

### Error: "go: command not found"

- Instala Go desde https://go.dev/dl/
- Reinicia tu terminal/VS Code
- Verifica: `go version`

---

## 🎯 Flujo de Trabajo Completo

```powershell
# 1. Iniciar servicios
.\scripts\start-services.ps1
.\scripts\init-db.ps1

# 2. En una nueva terminal: Iniciar backend
cd backend
go run cmd/server/main.go

# 3. En otra terminal: Sincronizar datos
Start-Sleep -Seconds 3
Invoke-RestMethod -Method POST -Uri "http://localhost:8080/api/sync"

# 4. Verificar
curl http://localhost:8080/api/stocks
```

---

## 📚 Recursos Adicionales

- [CockroachDB Docs](https://www.cockroachlabs.com/docs/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [Go Documentation](https://go.dev/doc/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 🤝 ¿Necesitas Ayuda?

Si encuentras algún problema:

1. Revisa los logs: `.\scripts\view-logs.ps1`
2. Verifica que Docker esté corriendo
3. Asegúrate de que todos los puertos estén disponibles
4. Revisa el archivo `.env`

¡Buena suerte! 🚀
