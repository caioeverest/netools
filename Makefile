#!make

# Install netools
all: lint
    go build -ldflags '-w -s' -a -o ./bin/netools ./pkg/cmd/cmd.go; \
    chmod +x ./bin/netools; \
    echo 'netools is installed in the bin directory'

lint:
    gofmt -s -w .; \
    chmod +x ./bin/golangci-lint; \
    ./bin/golangci-lint run; \

clean:
    rm -rf ./bin/netools;
