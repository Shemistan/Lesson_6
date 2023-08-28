test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"


mock-gen:
	mockgen -source=internal/storage/storage.go \
	-destination=internal/storage/mocks/mock_storage.go