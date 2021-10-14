#!make

all:
	# Install all networking tools

ci-lint:
	gofmt -s -w .; \
	chmod +x ./bin/golangci-lint; \
	./bin/golangci-lint run; \

clean:
	rm -rf ./bin/example_cli;