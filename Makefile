test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

mock:
	mockgen -source=./storage/storage.go -destination=./storage/mock/storage_mock.go
	mockgen -source=./service/userService.go -destination=./service/mock/userService_mock.go