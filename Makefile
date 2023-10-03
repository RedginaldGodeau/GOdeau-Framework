pwd ?= $(CURDIR)

.PHONY: ci
ci:
	@docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.54.2 golangci-lint run -v

.PHONY: run
run:
	@go run .

.PHONY: build
build:
	@go build -o $(pwd)/var/


