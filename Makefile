NAME := ship-manage-local

VERSION := v0.0.3

# 目标输出目录
DIST_FOLDER := dist

# 版本构建目录
RELEASE_FOLDER := resources

# proto 文件
API_PROTO_FILES=$(shell find api -name *.proto)


build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./dist/${NAME} ./cmd/${NAME}


container:
	docker build -t ${NAME}:${VERSION} -f ./deploy/Dockerfile .;


clean:
	-rm -rf ${DIST_FOLDER}
	-rm -rf api/**/*.go
	-go clean
	-go clean -cache
	-docker rmi $(docker images | grep "<none>" | awk '{print $3}')
