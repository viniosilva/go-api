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
	rm ./**/*_mock.go
	go generate ./...

swagger:
	swag init