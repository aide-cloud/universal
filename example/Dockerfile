FROM golang:1.18 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN make build

FROM alpine:latest

LABEL maintainer="Docker <biao.hu 1058165620@qq.com>"

WORKDIR /app

COPY --from=builder /app/bin ./bin
COPY --from=builder /app/configs ./configs

RUN apk --no-cache add ca-certificates

# 增加配置文件volume
VOLUME ["/app/configs"]

EXPOSE 18000

ENTRYPOINT ["./bin/cmd"]