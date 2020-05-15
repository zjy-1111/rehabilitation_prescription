.PHONY: build docker clean tool lint help scp

all: build

build:
	# 在对应命令前加上 @，可指定该命令不被打印到标准输出上
	#@go build -v .
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./output/rehabilitation .
	cp -r ./conf Dockerfile ./runtime ./output

docker:
    # 构建docker镜像
	docker build -t rehabilitation-scratch ./output
	# 保存docker镜像文件到本地
	docker save rehabilitation-scratch -o rehabilitation.tar
	# scp远程复制镜像文件到服务器
	scp ./rehabilitation.tar develop@47.100.213.205:~
	# 删除本地镜像文件
	rm -rf rehabilitation.tar
	# 删除远程容器镜像
	ssh develop@47.100.213.205 "docker rmi -f rehabilitation-scratch"
	# 加载新的镜像
	ssh develop@47.100.213.205 "docker load < rehabilitation.tar"
	# 杀掉之前的容器进程
	ssh develop@47.100.213.205 "docker rm -f `docker ps -a | grep "./rehabilitation" | awk '{print $1}'`"
	# 启动进程
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

