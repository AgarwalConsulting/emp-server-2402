ci-test:
	go generate ./...
	go test -v ./...