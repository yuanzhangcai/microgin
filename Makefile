
default: build

USER=`whoami`
GIT_TAG=`git describe --tags`
GIT_COMMIT=`git rev-parse HEAD`
BUILD_TIME=`date +%FT%T`
VERSION=`git rev-list --tags --max-count=1 | xargs git describe --tags`
LDFLAGS="-X main._commit=${GIT_COMMIT} -X main._buildtime=${BUILD_TIME} -X main._buildby=${USER} -X main._version=${VERSION}"

# 安装gowacth，执行该命令会给go.mod添加gowatch相关的依赖包，建议到工作目录外手动执行 go get
install:
	go get github.com/silenceper/gowatch

# 热编译启动程序 默认指定运行环境为本地运行环境，配置详见gowatch.yml
watch:
	gowatch

# 编译本地开发运行环境程序
dev:
	go build -ldflags ${LDFLAGS}" -X main._env=dev" -race

# 编译测试运行环境程序
test:
	go build -ldflags ${LDFLAGS}" -X main._env=test" -race

# 编译正式运行环境程序
prod: 
	go build -ldflags ${LDFLAGS}" -X main._env=prod" 

.PHONY: install watch dev test prod