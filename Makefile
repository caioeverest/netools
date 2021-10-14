#!make

all:
	# Install all networking tools

install-example-cli:
	go build -ldflags '-w -s' -a -o ./bin/example_cli ./example_cli/main.go; \
    chmod +x ./bin/example_cli; \

ci-lint:
	chmod +x ./bin/golangci-lint; \
	./bin/golangci-lint run; \

clean:
	rm -rf ./bin/example_cli;