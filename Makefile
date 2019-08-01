.PHONY: build clean tool lint help

all: build

build:
	# 在对应命令前加上 @，可指定该命令不被打印到标准输出上
	@go build -v .

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf rehabilitation_prescription
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
