.PHONY: run
run:
	docker-compose up -d
	go run cmd/server/main.go

test:
	go test -v ./...