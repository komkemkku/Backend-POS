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

# Auto deploy (commit + push)
auto-deploy:
	@echo "🚀 Running auto-deploy..."
	./auto-deploy.sh

# Deploy alias
deploy: auto-deploy

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
	@echo "  make deploy      - Auto commit & push changes"
	@echo "  make auto-deploy - Same as deploy"
	@echo "  make setup       - Setup development environment"
	@echo "  make help        - Show this help"

# Default target
default: help
