rpc-service:
	env GO111MODULE=on go build $(LDFLAGS)
.PHONY: rpc-service

clean:
	rm rpc-service

test:
	go test -v ./...

lint:
	golangci-lint run ./...