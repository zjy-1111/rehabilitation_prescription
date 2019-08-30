.PHONY: build docker clean tool lint help scp

all: build

build:
	# 在对应命令前加上 @，可指定该命令不被打印到标准输出上
	#@go build -v .
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./output/rehabilitation .
	cp -r ./conf Dockerfile ./runtime ./output

docker:
	docker build -t rehabilitation-scratch ./output
	docker save rehabilitation-scratch -o rehabilitation.tar
	scp ./rehabilitation.tar develop@47.100.213.205:~
	rm -rf rehabilitation.tar
	ssh develop@47.100.213.205 "docker rmi -f rehabilitation-scratch"
	ssh develop@47.100.213.205 "docker load < rehabilitation.tar"
	ssh develop@47.100.213.205 "docker rm -f `docker ps -a | grep "./rehabilitation" | awk '{print $1}'`"
	ssh develop@47.100.213.205 "docker run --link mysql:mysql -p 8000:8000 -d rehabilitation-scratch"

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf rehabilitation
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"

