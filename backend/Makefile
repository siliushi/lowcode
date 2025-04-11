ReleaseVersion = $(shell git describe --tags)
BuildTS = $(shell date '+%Y-%m-%d %H:%M:%S')
GitHash = $(shell git rev-parse HEAD)
GitBranch = $(shell git rev-parse --abbrev-ref HEAD)
GoVersion = $(shell go version)

LDFLAGS += -X 'lowcode/internal/version.ReleaseVersion=${ReleaseVersion}'
LDFLAGS += -X 'lowcode/internal/version.BuildTS=${BuildTS}'
LDFLAGS += -X 'lowcode/internal/version.GitHash=${GitHash}'
LDFLAGS += -X 'lowcode/internal/version.GitBranch=${GitBranch}'
LDFLAGS += -X 'lowcode/internal/version.GoVersion=${GoVersion}'

.PHONY: env genapi gensql clean lint build

all: env genapi gensql clean lint build

env:
	@echo "=========install goimports ==========="
	GOPROXY=https://goproxy.cn/,direct go install mvdan.cc/gofumpt@latest
	GOPROXY=https://goproxy.cn/,direct go install -v github.com/incu6us/goimports-reviser/v3@latest
	@echo "=========install golangci-lint for static code check========="
	GOPROXY=https://goproxy.cn/,direct go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "=========install goctl for generate api project========="
	GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
	@echo "=========install goctl-swagger for generate swagger json========="
	GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest
	@echo "=========install gentool for generate sql model========="
	GOPROXY=https://goproxy.cn/,direct go install gorm.io/gen/tools/gentool@v0.3.18


gensql:
	gentool -dsn "lowcode:qaz123!@#@tcp(127.0.0.1:3306)/lowcodedb?charset=utf8mb4&parseTime=True&loc=Local" -outPath "./internal/dao/query"

genapi:
	goctl api format -dir api
	goctl api go -api api/main.api -dir .
	goctl api plugin -plugin goctl-swagger="swagger -filename lowcode.json -host 127.0.0.1:8888 -basepath /" -api api/main.api -dir ./internal/handler/swagger

build:
	go mod tidy
	gofumpt -l -w .
	goimports-reviser -rm-unused -format ./...
	CGO_ENABLED=0 go build  -ldflags "$(LDFLAGS)" -v .

lint:
	golangci-lint run ./...

test:
	go test ./...

clean:
	go clean -i .

run:
	./lowcode
