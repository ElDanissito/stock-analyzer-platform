#!/bin/bash

# Script para inicializar la base de datos CockroachDB

echo "🚀 Inicializando base de datos CockroachDB..."

# Esperar a que CockroachDB esté listo
echo "⏳ Esperando a que CockroachDB esté listo..."
sleep 5

# Crear la base de datos
echo "📦 Creando base de datos 'stocks'..."
docker exec -it stock-analyzer-cockroachdb ./cockroach sql --insecure -e "CREATE DATABASE IF NOT EXISTS stocks;"

# Verificar que la base de datos fue creada
echo "✅ Verificando base de datos..."
docker exec -it stock-analyzer-cockroachdb ./cockroach sql --insecure -e "SHOW DATABASES;"

echo "✨ Base de datos inicializada correctamente!"
echo ""
echo "📊 Puedes acceder al Admin UI en: http://localhost:8080"
echo "🔌 Connection string: postgresql://root@localhost:26257/stocks?sslmode=disable"
