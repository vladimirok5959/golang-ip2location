test:
	go test ./...

bench:
	go test ./... -run=NONE -bench=. -benchmem

staticcheck:
	staticcheck ./...

lint:
	golangci-lint run

tidy:
	go mod tidy

.PHONY: test bench staticcheck tidy
