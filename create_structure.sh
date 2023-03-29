#!/bin/bash

echo "Creando estructura de carpetas y archivos..."

# Crea los directorios
mkdir -p miweb/cmd/web
mkdir -p miweb/internal/config
mkdir -p miweb/internal/handlers
mkdir -p miweb/internal/models
mkdir -p miweb/internal/repositories
mkdir -p miweb/internal/services
mkdir -p miweb/static
mkdir -p miweb/templates

# Crea los archivos
touch miweb/cmd/web/main.go
touch miweb/internal/config/config.go
touch miweb/internal/handlers/handlers.go
touch miweb/internal/models/models.go
touch miweb/internal/repositories/repositories.go
touch miweb/internal/services/services.go

echo "Estructura de carpetas y archivos creada."
