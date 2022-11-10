VERSION=$(shell git describe --tags --always)
IMAGE_NAME=$(shell basename $(shell pwd))
DOCKER_IMAGE="$(IMAGE_NAME)"

# 是否存在bin目录,如果不存在则创建
ifeq (,$(wildcard ./bin))
$(shell mkdir "bin")
endif

.PHONY: init
# 初始化go运行依赖
init:
	go mod tidy

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: dev
# 本地开发
dev:
	go run -ldflags "-X main.Version=$(VERSION)" cmd/main.go

.PHONY: test
# 单元测试
test:
	go test -v ./...

.PHONY: image
# 构建docker镜像
image:
	@echo "build image: $(DOCKER_IMAGE):$(VERSION)";
	@echo "";
	@docker build -t "$(DOCKER_IMAGE):latest" .;
	@docker system prune -f;
	@echo "--------------------------------------------------";

# 帮助
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help