# Change this to your own Go module
GO_MODULE := protobuf-go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf ./protogen

.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	./proto/basic/*.proto \


.PHONY: build
build: clean protoc tidy

.PHONY: run
run:
	go run main.go

.PHONY: execute
execute: build run

.PHONY: protoc-validate
protoc-validate:
	protoc --validate_out="lang=go:./generated" --go_opt=module=${GO_MODULE} --go_out=. ./proto/car/*.proto
