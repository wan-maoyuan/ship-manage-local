# 程序编译镜像
FROM golang:1.23 AS build

ENV GO111MODULE=on                               \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /go/release

ADD . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./dist/ship-manage-local ./cmd/ship-manage-local

# 构建镜像
FROM debian:stable-slim AS prod

ENV TZ=Asia/Shanghai                                        \
    DEBIAN_FRONTEND=noninteractive                          \
    GIN_MODE=release                                        \
    ADDR=":8888"

RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev && rm -rf /var/lib/apt/lists/*

# 在build阶段复制时区到
COPY --from=build /go/release/dist/ship-manage-local    /ship-manage-local
COPY --from=build /go/release/resources/db_files        /resources/db_files

CMD [ "/ship-manage-local"]