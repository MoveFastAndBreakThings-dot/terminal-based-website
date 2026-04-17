BINARY   := ssh-portfolio
PORT     := 23234
SSH_HOST := localhost

.PHONY: all build run test clean docker-up docker-down docker-logs connect fmt vet lint ship

all: build

## ── Ship (build + test + deploy) ─────────────────────────────────────────────

ship: fmt vet build test docker-up
	@echo ""
	@echo "  ✓ fmt · vet · build · test · deployed on :$(PORT)"
	@echo ""

## ── Build ─────────────────────────────────────────────────────────────────────

build:
	go build -ldflags="-s -w" -o $(BINARY) .

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BINARY)-linux .

## ── Dev ───────────────────────────────────────────────────────────────────────

run:
	go run .

connect:
	ssh -p $(PORT) $(SSH_HOST)

## ── Test ──────────────────────────────────────────────────────────────────────

test:
	go test ./tests/... -v

test-short:
	go test ./tests/... -short

test-race:
	go test ./tests/... -race

cover:
	go test ./tests/... -coverprofile=coverage.out
	go tool cover -html=coverage.out

## ── Quality ───────────────────────────────────────────────────────────────────

fmt:
	gofmt -w .

vet:
	go vet ./...

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

## ── Docker ────────────────────────────────────────────────────────────────────

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

docker-restart:
	docker compose restart

## ── Clean ─────────────────────────────────────────────────────────────────────

clean:
	rm -f $(BINARY) $(BINARY)-linux coverage.out

## ── Help ──────────────────────────────────────────────────────────────────────

help:
	@echo ""
	@echo "  build          compile binary"
	@echo "  build-linux    cross-compile for linux/amd64"
	@echo "  run            go run . (dev server on :$(PORT))"
	@echo "  connect        ssh -p $(PORT) $(SSH_HOST)"
	@echo "  test           run all tests verbose"
	@echo "  test-short     run tests (skip slow)"
	@echo "  test-race      run tests with race detector"
	@echo "  cover          generate HTML coverage report"
	@echo "  fmt            gofmt all files"
	@echo "  vet            go vet"
	@echo "  lint           golangci-lint"
	@echo "  tidy           go mod tidy"
	@echo "  docker-up      build + start container"
	@echo "  docker-down    stop container"
	@echo "  docker-logs    tail container logs"
	@echo "  docker-restart restart container"
	@echo "  clean          remove build artifacts"
	@echo "  ship           fmt + vet + build + test + docker-up (one command)"
	@echo ""
