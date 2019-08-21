BIN := mini_spider

.PHONY: all
all: build

.PHONY: build-output
build-output: clean
	@mkdir -p build_output/

.PHONY: build
build: build-output
	@go build -o ./${BIN} 

.PHONY: clean
clean:
	@rm -rf build_output