.PHONY: run
run:
	docker-compose up -d
	go run server/main.go

test:
	go test -v ./...