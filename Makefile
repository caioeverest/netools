#!make

all:
	# Install all networking tools

install-example-cli:
	go build -ldflags '-w -s' -a -o ./bin/example_cli ./src/example_cli/main.go; \
    chmod +x ./bin/example_cli; \

install-subnet-calculator:
	go build -ldflags '-w -s' -a -o ./bin/subnet_calculator ./src/subnet_calculator/main.go; \
    chmod +x ./bin/subnet_calculator; \

ci-lint:
	chmod +x ./bin/golangci-lint; \
	./bin/golangci-lint run; \

clean:
	rm -rf ./bin/example_cli; \
	rm -rf ./bin/subnet_calculator; \
