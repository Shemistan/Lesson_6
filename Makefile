mock-gen:
	mockgen -source=internal/storage/storage.go \
	-destination=internal/storage/mocks/mock_storage.go
mock-gen1:
	mockgen -source=internal/service/service.go \
	-destination=internal/service/mocks/mock_service.go
test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"


