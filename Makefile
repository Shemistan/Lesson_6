mock-gen:
	mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock_service.go

test-run:
	go test -v ./...

	