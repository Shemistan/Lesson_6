test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"


mock:
	mockgen -source=internal/storage/storage.go -destination=internal/storage/mock/storage_mock.go
	mockgen -source=internal/service/service.go -destination=internal/service/mock/service_mock.go
