infra/up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

run:
	go run main.go

.PHONY: test
test:
	go test -v ./...

test/cov:
	go test -v ./... -cov

update:
	go mod tidy

mock:
	find . -name "*_mock.go" -type f -delete
	go generate ./...

swagger:
	swag init