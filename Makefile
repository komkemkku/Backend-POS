# Makefile สำหรับ Backend POS
# ใช้คำสั่ง: make <target>

.PHONY: dev build deploy auto-deploy test clean migrate

# Development
dev:
	@echo "🚀 Starting development server..."
	go run main.go

# Build application
build:
	@echo "🔨 Building application..."
	go build -o pos-server .

# Run migration
migrate:
	@echo "📊 Running database migration..."
	go run cmd/migrateCmd.go

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	go mod tidy
	go mod download

# Test application
test:
	@echo "🧪 Running tests..."
	go test ./...

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -f pos-server
	rm -f main
	go clean

# Deep clean (Windows)
cleanup:
	@echo "🧹 Running deep cleanup..."
	powershell -ExecutionPolicy Bypass -File ./cleanup.ps1

# Auto deploy (commit + push)
auto-deploy:
	@echo "🚀 Running auto-deploy..."
	./auto-deploy.sh

# Auto deploy for Windows (PowerShell)
auto-deploy-win:
	@echo "🚀 Running auto-deploy (Windows)..."
	powershell -ExecutionPolicy Bypass -File ./auto-deploy.ps1

# Auto watch and deploy (Windows)
watch-deploy:
	@echo "🔍 Starting auto-watch for changes..."
	powershell -ExecutionPolicy Bypass -File ./auto-watch.ps1

# Deploy alias
deploy: auto-deploy

# Deploy for Windows
deploy-win: auto-deploy-win

# Setup development environment
setup:
	@echo "⚙️ Setting up development environment..."
	cp .env.example .env
	@echo "✅ Please edit .env file with your database credentials"
	@echo "📊 Run 'make migrate' after setting up database"
	@echo "🚀 Run 'make dev' to start development server"

# Show help
help:
	@echo "📋 Available commands:"
	@echo "  make dev         - Start development server"
	@echo "  make build       - Build the application"
	@echo "  make migrate     - Run database migration"
	@echo "  make deps        - Install dependencies"
	@echo "  make test        - Run tests"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make cleanup     - Deep clean (remove all unnecessary files)"
	@echo "  make deploy      - Auto commit & push changes (Linux/Mac)"
	@echo "  make deploy-win  - Auto commit & push changes (Windows)"
	@echo "  make watch-deploy- Auto-watch files and deploy on changes"
	@echo "  make auto-deploy - Same as deploy"
	@echo "  make auto-deploy-win - Same as deploy-win"
	@echo "  make setup       - Setup development environment"
	@echo "  make help        - Show this help"

# Default target
default: help
