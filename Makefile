build:
	docker-compose up -d --build

test:
	docker-compose exec app sh -c "go test ./... -cover"

.PHONY: build test