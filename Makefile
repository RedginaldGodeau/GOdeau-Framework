pwd ?= $(CURDIR)

.PHONY: ci
ci:
	@docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v

.PHONY: run
run:
	@go run .

build: windows-build linux-build
windows-build:
	@set GOOS=windows
	@go build -o $(pwd)/var/app.exe
linux-build:
	@set GOOS=windows
	@go build -o $(pwd)/var/app

.PHONY: start
start:
	docker compose up -d

.PHONY: start
down:
	docker compose down