.PHONY: run
run:
	docker-compose up -d
	go run cmd/server/main.go

test:
	docker-compose up -d
	go test -v ./...
	docker-compose down